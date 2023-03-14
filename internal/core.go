package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"go.uber.org/zap"
)

// GitCommit,BuiltDate are set at build-time
var Version = "v0.0.1-SNAPSHOT"
var GitCommit = "54a8d74ea3cf6fdcadfac10ee4a4f2553d4562f6q"
var BuiltDate = "Thu Jan  1 01:00:00 CET 1970" //date -r 0 (Mac), date -d @0 (Linux)

func Print_Version() {
	fmt.Printf("Client: CUC - Community\nVersion: \t%s\nGit commit: \t%s\nBuilt: \t\t%s\n", Version, GitCommit, BuiltDate)
}

func Check_URL(URL, musicFile string, backoff, httpCode int, loop bool, logger *zap.Logger, ctx context.Context) {
	var attempt int = 1

	ctx, cancel := context.WithCancel(ctx)

	client := &http.Client{}

	// #nosec [G304] [-- Acceptable risk, for the CWE-22]
	f, err := os.Open(musicFile)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	// ../../../../go/pkg/mod/github.com/hajimehoshi/oto@v1.0.1/context.go:69:12: undefined: newDriver
	// To fix this error, we need to enable CGO_ENABLED=1
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}

	//https://github.com/faiface/beep/wiki/To-buffer,-or-not-to-buffer,-that-is-the-question
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	err = streamer.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Graceful shutdown goroutine
	go func(context.CancelFunc) {
		sigquit := make(chan os.Signal, 1)
		// os.Kill can't be caught https://groups.google.com/g/golang-nuts/c/t2u-RkKbJdU
		// POSIX spec: signal can be caught except SIGKILL/SIGSTOP signals
		// Ctrl-c (usually) sends the SIGINT signal, not SIGKILL
		// syscall.SIGTERM usual signal for termination
		// and default one for docker containers, which is also used by kubernetes
		signal.Notify(sigquit, os.Interrupt, os.Kill, syscall.SIGTERM)
		sig := <-sigquit

		logger.Info("Caught the following signal", zap.String("signal", sig.String()))
		cancel()
	}(cancel)

	duration := time.Duration(backoff) * time.Second
	ticker := time.NewTicker(duration)
	done := make(chan bool)
	for {
		select {
		case <-ticker.C:
			req, err1 := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
			if err != nil {
				fmt.Println(err1)
				return
			}
			resp, err2 := client.Do(req)
			if err != nil {
				fmt.Println(err2)
				return
			}
			// To avoid this error panic: runtime error: invalid memory address or nil pointer dereference
			// [signal SIGSEGV: segmentation violation code=0x1 addr=0x10 pc=0x1002f8e34]
			/*if resp == nil {
				graceful_shutdown(logger, ctx)
				defer resp.Body.Close()
				return
			}*/
			if resp.StatusCode == httpCode {
				logger.Info("It's a match!",
					zap.Int("attempt", attempt),
					zap.Int("statuscode", resp.StatusCode),
					zap.Duration("backoff", duration),
					zap.String("url", URL),
				)
				music := buffer.Streamer(0, buffer.Len())
				speaker.Play(beep.Seq(music, beep.Callback(func() {
					done <- true
				})))
				<-done
				if !loop {
					defer func() {
						if err := resp.Body.Close(); err != nil {
							logger.Info("Error closing the http.NewRequest.body:", zap.Error(err))
						}
					}()
					graceful_shutdown(logger, ctx)
					return
				}
			} else {
				logger.Info("Unmatch status code",
					zap.Int("attempt", attempt),
					zap.Int("statuscode", resp.StatusCode),
					zap.Duration("backoff", duration),
					zap.String("url", URL),
				)
			}
			defer func() {
				if err := resp.Body.Close(); err != nil {
					logger.Info("Error closing the http.NewRequest.body:", zap.Error(err))
				}
			}()
			attempt++
		case <-ctx.Done():
			graceful_shutdown(logger, ctx)
			return
		}
	}
}

func graceful_shutdown(logger *zap.Logger, ctx context.Context) {
	if ctx.Err() == nil {
		logger.Info("Graceful shutdown..")
	} else {
		logger.Info("Graceful shutdown..",
			zap.String("ctx.err", ctx.Err().Error()),
		)
	}

}
