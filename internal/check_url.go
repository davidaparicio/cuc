/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
*/
package internal

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func CheckURL(url, musicFile string, backoff, httpCode int, loop bool, logger *zap.Logger, cmd *cobra.Command) {
	var attempt int = 1
	ctx, cancel := context.WithCancel(cmd.Root().Context())

	// Graceful shutdown goroutine
	go func(context.CancelFunc) {
		sigquit := make(chan os.Signal, 1)
		// POSIX: Ctrl-c (usually) sends the SIGINT signal
		// syscall.SIGTERM usual signal for termination
		// and default one for docker containers, which is also used by kubernetes
		signal.Notify(sigquit, os.Interrupt, syscall.SIGTERM)
		sig := <-sigquit

		logger.Info("Caught the following signal", zap.String("signal", sig.String()))
		cancel()
	}(cancel)

	buffer, err := prepareMusic(musicFile, logger)
	if err != nil {
		logger.Fatal("Unexpected error while executing command:",
			zap.String("prepareMusic err", err.Error()))
		return
	}

	duration := time.Duration(backoff) * time.Second
	ticker := time.NewTicker(duration)
	done := make(chan bool)
	breaking := false
	for !breaking {
		select {
		case <-ticker.C:
			code, err := check(url, logger, ctx)
			if err != nil {
				logger.Fatal("Unexpected error while executing command:",
					zap.String("check err", err.Error()))
				return
			}
			if code == httpCode {
				logger.Info("It's a match!",
					zap.Int("attempt", attempt),
					zap.Int("statuscode", code),
					zap.Duration("backoff", duration),
					zap.String("url", url),
				)
				music := buffer.Streamer(0, buffer.Len())
				speaker.Play(beep.Seq(music, beep.Callback(func() {
					done <- true
				})))
				<-done
				if !loop {
					breaking = true
				}
			} else {
				logger.Info("Unmatch status code",
					zap.Int("attempt", attempt),
					zap.Int("statuscode", code),
					zap.Duration("backoff", duration),
					zap.String("url", url),
				)
			}
			attempt++
		case <-ctx.Done():
			breaking = true
		}
	}
	gracefulShutdown(logger, ctx)
}

func check(url string, logger *zap.Logger, ctx context.Context) (statuscode int, err error) {
	var resp *http.Response
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		logger.Fatal("Unexpected error while executing command:",
			zap.String("http.NewRequestWithContext err", err.Error()))
		return http.StatusInternalServerError, err
	}
	resp, err = client.Do(req)
	if err != nil {
		logger.Fatal("Unexpected error while executing command:",
			zap.String("client.Do err", err.Error()))
		return http.StatusInternalServerError, err
	}
	// To avoid this error panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x10 pc=0x1002f8e34]
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.Info("Error closing the http.NewRequest.body:", zap.Error(err))
		}
	}()
	return resp.StatusCode, nil
}

func gracefulShutdown(logger *zap.Logger, ctx context.Context) {
	if ctx.Err() == nil {
		logger.Info("Graceful shutdown..")
	} else {
		logger.Info("Graceful shutdown..",
			zap.String("ctx.err", ctx.Err().Error()),
		)
	}
}
