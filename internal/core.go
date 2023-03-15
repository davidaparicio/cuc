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
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	SpeakerRate = 10
)

// Version GitCommit BuiltDate are set at build-time
var Version = "v0.0.1-SNAPSHOT"
var GitCommit = "54a8d74ea3cf6fdcadfac10ee4a4f2553d4562f6q"
var BuildDate = "Thu Jan  1 01:00:00 CET 1970" // date -r 0 (Mac), date -d @0 (Linux)

func PrintVersion(cmd *cobra.Command) {
	cmd.Printf("Client: CUC - Community\nVersion: \t%s\nGit commit: \t%s\nBuilt: \t\t%s\n", Version, GitCommit, BuildDate)
}

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

	var resp *http.Response
	client := &http.Client{}
	duration := time.Duration(backoff) * time.Second
	ticker := time.NewTicker(duration)
	done := make(chan bool)
	breaking := false
	for !breaking {
		select {
		case <-ticker.C:
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
			if err != nil {
				logger.Fatal("Unexpected error while executing command:",
					zap.String("http.NewRequestWithContext err", err.Error()))
				return
			}
			resp, err = client.Do(req)
			if err != nil {
				logger.Fatal("Unexpected error while executing command:",
					zap.String("client.Do err", err.Error()))
				return
			}
			// To avoid this error panic: runtime error: invalid memory address or nil pointer dereference
			// [signal SIGSEGV: segmentation violation code=0x1 addr=0x10 pc=0x1002f8e34]
			if resp.StatusCode == httpCode {
				logger.Info("It's a match!",
					zap.Int("attempt", attempt),
					zap.Int("statuscode", resp.StatusCode),
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
					zap.Int("statuscode", resp.StatusCode),
					zap.Duration("backoff", duration),
					zap.String("url", url),
				)
			}
			attempt++
		case <-ctx.Done():
			breaking = true
		}
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.Info("Error closing the http.NewRequest.body:", zap.Error(err))
		}
	}()
	gracefulShutdown(logger, ctx)
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

func prepareMusic(musicFile string, logger *zap.Logger) (buffer *beep.Buffer, err error) {
	// #nosec [G304] [-- Acceptable risk, for the CWE-22]
	f, err := os.Open(musicFile)
	if err != nil {
		logger.Warn("Not possible to open the file", zap.String("os.Open err", err.Error()))
		return nil, err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		logger.Warn("Not possible to decode the MP3 file", zap.String("mp3.Decode err", err.Error()))
		return nil, err
	}

	// ../../../../go/pkg/mod/github.com/hajimehoshi/oto@v1.0.1/context.go:69:12: undefined: newDriver
	// To fix this error, we need to enable CGO_ENABLED=1
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/SpeakerRate))
	if err != nil {
		logger.Warn("Not possible to init the speaker", zap.String("speaker.Init err", err.Error()))
		return nil, err
	}

	//https://github.com/faiface/beep/wiki/To-buffer,-or-not-to-buffer,-that-is-the-question
	buffer = beep.NewBuffer(format)
	buffer.Append(streamer)
	err = streamer.Close()
	if err != nil {
		logger.Warn("Not possible to close the streamer", zap.String("streamer.Close err", err.Error()))
		return nil, err
	}
	return buffer, nil
}
