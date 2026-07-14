//go:build !cgo && !windows

/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
*/
package internal

import (
	"os"

	"go.uber.org/zap"
)

type silentPlayer struct {
	logger *zap.Logger
}

// prepareMusic only validates that the MP3 file is readable: the audio
// backend needs CGO on this platform and this binary was built with
// CGO_ENABLED=0, so a silent player is returned instead.
func prepareMusic(musicFile string, logger *zap.Logger) (player, error) {
	// #nosec [G304] [-- Acceptable risk, for the CWE-22]
	f, err := os.Open(musicFile)
	if err != nil {
		logger.Warn("Not possible to open the file", zap.String("os.Open err", err.Error()))
		return nil, err
	}
	if err := f.Close(); err != nil {
		logger.Warn("Not possible to close the file", zap.String("f.Close err", err.Error()))
		return nil, err
	}
	logger.Warn("This build has no audio support (compiled with CGO_ENABLED=0): the cheer will be silent")
	return &silentPlayer{logger: logger}, nil
}

func (p *silentPlayer) Play() {
	p.logger.Warn("Audio support unavailable in this build, skipping the jingle")
}
