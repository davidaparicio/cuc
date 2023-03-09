/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
*/
package cmd

import (
	"os"

	"github.com/davidaparicio/cuc/internal"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	//sugar     *zap.SugaredLogger
	verbose   bool
	URL       string
	musicFile string
	backoff   int
	httpCode  int
	//cfgFile string
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
		/*URL, err1 := cmd.Flags().GetString("URL")
		musicFile, err2 := cmd.Flags().GetString("musicFile")
		backoff, err3 := cmd.Flags().GetInt("seconds")
		want, err4 := cmd.Flags().GetInt("httpCode")

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			log.Fatal("Can't parse args")
		}*/

		internal.Check_URL(URL, musicFile, backoff, httpCode, false, logger, cmd.Root().Context())
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&URL, "URL", "u", "https://www.example.com/", "Webpage to check")
	rootCmd.PersistentFlags().StringVarP(&musicFile, "musicFile", "f", "./assets/mp3/ubuntu_desktop_login.mp3", "MP3 file to play if the check is successful")
	rootCmd.PersistentFlags().IntVarP(&backoff, "seconds", "s", 30, "Backoff in seconds")
	rootCmd.PersistentFlags().IntVarP(&httpCode, "httpCode", "c", 200, "HTTP Status Code from 100 to 511")
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cuc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enables debug logging")

	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	//rand.Seed(int64(time.Now().Nanosecond()))

	zapOptions := []zap.Option{
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	}
	if !verbose {
		zapOptions = append(zapOptions,
			zap.IncreaseLevel(zap.LevelEnablerFunc(func(l zapcore.Level) bool { return l != zapcore.DebugLevel })),
		)
	}
	//l, _ := zap.NewDevelopment(zapOptions...)
	l, _ := zap.NewProduction(zapOptions...)
	defer l.Sync() // flushes buffer, if any
	// L returns the global Logger, which can be reconfigured with ReplaceGlobals.
	// It's safe for concurrent use.
	undo := zap.ReplaceGlobals(l)
	defer undo()
	logger = zap.L()
	//logger.Info("replaced zap's global loggers")

	/*if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cuc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cuc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}*/
}

/*func logError(logger *zap.Logger, err error) error {
	if err != nil {
		logger.Error("Error running command", zap.Error(err))
	}
	return err
}*/
