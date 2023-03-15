/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
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
		internal.PrintVersion(cmd)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
