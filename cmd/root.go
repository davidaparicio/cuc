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

var (
	logger *zap.Logger
	// sugar     *zap.SugaredLogger
	verbose   bool
	URL       string
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
			internal.PrintVersion()
		} else {
			internal.CheckURL(URL, musicFile, timeout, httpCode, false, logger, cmd.Root().Context())
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

func init() {
	rootCmd.PersistentFlags().StringVarP(&URL, "URL", "u", "https://www.example.com/", "Webpage to check")
	rootCmd.PersistentFlags().StringVarP(&musicFile,
		"musicFile", "f", "./assets/mp3/ubuntu_desktop_login.mp3", "MP3 file to play if the check is successful")
	rootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 1, "Timeout in seconds")
	rootCmd.PersistentFlags().IntVarP(&httpCode, "httpCode", "c", 200, "HTTP Status Code from 100 to 511")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "d", false, "Enables debug logging")
	rootCmd.PersistentFlags().BoolVarP(&printVersion, "version", "v", false, "Print the version")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	zapOptions := []zap.Option{
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	}
	if !verbose {
		zapOptions = append(zapOptions,
			zap.IncreaseLevel(zap.LevelEnablerFunc(func(l zapcore.Level) bool { return l != zapcore.DebugLevel })),
		)
	}
	l, _ := zap.NewProduction(zapOptions...)
	defer func() {
		if err := l.Sync(); err != nil { // flushes buffer, if any
			fmt.Println("Error during flushing all logger buffers (l.Sync())")
		}
	}()
	// L returns the global Logger, which can be reconfigured with ReplaceGlobals.
	// It's safe for concurrent use.
	undo := zap.ReplaceGlobals(l)
	defer undo()
	logger = zap.L()
}
