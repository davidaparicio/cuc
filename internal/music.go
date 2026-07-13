//go:build cgo || windows

/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
*/
package internal

import (
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
	"go.uber.org/zap"
)

const (
	SpeakerRate = 10
)

type beepPlayer struct {
	buffer *beep.Buffer
}

func prepareMusic(musicFile string, logger *zap.Logger) (player, error) {
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

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/SpeakerRate))
	if err != nil {
		logger.Warn("Not possible to init the speaker", zap.String("speaker.Init err", err.Error()))
		return nil, err
	}

	//https://github.com/faiface/beep/wiki/To-buffer,-or-not-to-buffer,-that-is-the-question
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	err = streamer.Close()
	if err != nil {
		logger.Warn("Not possible to close the streamer", zap.String("streamer.Close err", err.Error()))
		return nil, err
	}
	return &beepPlayer{buffer: buffer}, nil
}

func (p *beepPlayer) Play() {
	done := make(chan bool)
	music := p.buffer.Streamer(0, p.buffer.Len())
	speaker.Play(beep.Seq(music, beep.Callback(func() {
		done <- true
	})))
	<-done
}
