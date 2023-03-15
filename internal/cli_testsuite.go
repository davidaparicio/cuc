/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
*/
package internal

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

type Test struct {
	Title          string
	Cmd            *cobra.Command
	Args           []string
	ExpectedStdout string
}

func ExecuteSuite(t *testing.T, tests []Test) {
	t.Parallel()
	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			output := &bytes.Buffer{}
			test.Cmd.SetOut(output) // https://stackoverflow.com/a/66804032
			test.Cmd.SetArgs(test.Args)

			err := test.Cmd.Execute()

			if !IsNil(err) {
				t.Errorf("Unexpected error while executing command: %v", err)
			}
			if !IsEqualString(test.ExpectedStdout, output.String()) {
				t.Errorf("Unexpected Stdout\ngot:\n%s\nexpected:\n%s\n", output.String(), test.ExpectedStdout)
			}
		})
	}
}

func IsEqualString(s1, s2 string) bool {
	return s1 == s2
}

func IsNil(object any) bool {
	return object == nil
}
