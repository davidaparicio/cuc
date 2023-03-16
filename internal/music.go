/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
*/
package internal

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"go.uber.org/zap"
)

const (
	SpeakerRate = 10
)

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
