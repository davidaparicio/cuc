/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/davidaparicio/cuc/internal"
	"github.com/spf13/cobra"
)

// loopCmd represents the loop command
var loopCmd = &cobra.Command{
	Use:   "loop",
	Short: "Loop until desired HTTP status is reached",
	Long: `CUC is CLI tool to check various HTTP status.
It will loop until the desired HTTP status is reached.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Check_URL(URL, musicFile, backoff, httpCode, true, logger, cmd.Root().Context())
	},
}

func init() {
	rootCmd.AddCommand(loopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loopCmd.PersistentFlags().Int("seconds", 30, "Backoff in seconds")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
