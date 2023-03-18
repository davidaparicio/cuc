/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
*/
package cmd

import (
	"testing"

	"github.com/davidaparicio/cuc/internal"
)

var versionOutput = "Client: CUC - Community\n" +
	"Version: 	v0.0.1-SNAPSHOT\n" +
	"Git commit: 	54a8d74ea3cf6fdcadfac10ee4a4f2553d4562f6q\n" +
	"Built: 		Thu Jan  1 01:00:00 CET 1970\n"

func TestVersionCMD(t *testing.T) {
	// https://stackoverflow.com/a/23205902
	tests := []internal.Test{
		/*{
			Title:          "print short version",
			Cmd:            rootCmd,
			Args:           []string{"-v"},
			ExpectedStdout: versionOutput,
		},
		{
			Title:          "print long version",
			Cmd:            rootCmd,
			Args:           []string{"--version"},
			ExpectedStdout: versionOutput,
		},*/
		{
			Title:          "print command version",
			Cmd:            rootCmd,
			Args:           []string{"version"},
			ExpectedStdout: versionOutput,
		},
	}
	internal.ExecuteSuite(t, tests)
}
