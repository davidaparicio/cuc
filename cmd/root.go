/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/davidaparicio/cuc/internal"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Timeout  = 1
	HTTPCode = 200
)

var (
	logger *zap.Logger
	// sugar     *zap.SugaredLogger
	verbose   bool
	url       string
	musicFile string
	timeout   int
	backoff   int
	httpCode  int
	// cfgFile string
	printVersion bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cuc",
	Short: "A simple URL checker",
	Long: `A very simple CLI tool to check various HTTP status.
CUC can loop until the desired HTTP status is reached.
For example:

If a concert ticket webpage is available (200), or not found (404).`,
	Run: func(cmd *cobra.Command, args []string) {
		if printVersion {
			internal.PrintVersion(cmd)
		} else {
			internal.CheckURL(url, musicFile, timeout, httpCode, false, logger, cmd)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.ExecuteContext(rootCmd.Context())
	if err != nil {
		os.Exit(1)
	}
}

//nolint:gochecknoinits
func init() {
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "https://www.example.com/", "Webpage to check")
	rootCmd.PersistentFlags().StringVarP(&musicFile,
		"musicFile", "f", "./assets/mp3/ubuntu_desktop_login.mp3", "MP3 file to play if the check is successful")
	rootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", Timeout, "Timeout in seconds")
	rootCmd.PersistentFlags().IntVarP(&httpCode, "httpCode", "c", HTTPCode, "HTTP Status Code from 100 to 511")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "d", false, "Enables debug logging")
	rootCmd.PersistentFlags().BoolVarP(&printVersion, "version", "v", false, "Print the version")

	rootCmd.AddCommand(loopCmd)
	loopCmd.PersistentFlags().IntVarP(&backoff, "backoff", "b", Backoff, "Backoff in seconds")

	rootCmd.AddCommand(versionCmd)

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	zapOptions := []zap.Option{
		// skip the first caller, which is typically the logger's own function
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.FatalLevel),
	}
	if !verbose { // level "debug" is not allowed by existing core
		zapOptions = append(zapOptions, zap.IncreaseLevel(zapcore.FatalLevel))
	}
	l, err := zap.NewProduction(zapOptions...)
	if err != nil {
		fmt.Println("Error during setting the Uber Zap logging")
	}
	defer func() { // flushes buffer, if any
		if err := logger.Sync(); err != nil {
			fmt.Println("Error during flushing all logger buffers (l.Sync()): " + err.Error())
		}
	}()
	// L returns the global Logger, which can be reconfigured with ReplaceGlobals.
	// It's safe for concurrent use.
	undo := zap.ReplaceGlobals(l)
	logger = zap.L()
	undo()
}
