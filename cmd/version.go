/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/davidaparicio/cuc/internal"
	"github.com/spf13/cobra"
)

// versionCmd represents the loop command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  `Print the CUC version`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.PrintVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
