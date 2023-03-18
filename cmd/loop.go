/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
*/
package cmd

import (
	"github.com/davidaparicio/cuc/internal"
	"github.com/spf13/cobra"
)

const (
	Backoff = 30
)

// loopCmd represents the loop command
var loopCmd = &cobra.Command{
	Use:   "loop",
	Short: "Loop until desired HTTP status is reached",
	Long: `CUC is CLI tool to check various HTTP status.
It will loop until the desired HTTP status is reached.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.CheckURL(url, musicFile, timeout, httpCode, true, logger, cmd)
	},
}
