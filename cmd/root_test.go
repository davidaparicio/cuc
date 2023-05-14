/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
*/
package cmd

import (
	"testing"

	"github.com/davidaparicio/cuc/internal"
)

var helpOutput = `A very simple CLI tool to check various HTTP status.
CUC can loop until the desired HTTP status is reached.
For example:

If a concert ticket webpage is available (200), or not found (404).

Usage:
  cuc [flags]
  cuc [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  loop        Loop until desired HTTP status is reached
  version     Print the version

Flags:
  -h, --help               help for cuc
  -c, --httpCode int       HTTP Status Code from 100 to 511 (default 200)
  -f, --musicFile string   MP3 file to play if the check is successful (default "./assets/mp3/ubuntu_desktop_login.mp3")
  -t, --timeout int        Timeout in seconds (default 1)
  -u, --url string         Webpage to check (default "https://www.example.com/")
  -d, --verbose            Enables debug logging
  -v, --version            Print the version

Use "cuc [command] --help" for more information about a command.
`

func TestRootCMD(t *testing.T) {
	// https://stackoverflow.com/a/23205902
	tests := []internal.Test{
		{
			Title:          "print short help",
			Cmd:            rootCmd,
			Args:           []string{"-h"},
			ExpectedStdout: helpOutput,
		},
		{
			Title:          "print long help",
			Cmd:            rootCmd,
			Args:           []string{"--help"},
			ExpectedStdout: helpOutput,
		},
		{
			Title:          "print command help",
			Cmd:            rootCmd,
			Args:           []string{"help"},
			ExpectedStdout: helpOutput,
		},
		/*{
			Title:          "empty args",
			Cmd:            rootCmd,
			Args:           []string{""},
			ExpectedStdout: helpOutput,
		},*/
	}
	internal.ExecuteSuite(t, tests)
}
