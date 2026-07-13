/*
Copyright © 2023 David Aparicio david.aparicio@free.fr
*/
package internal

// player plays the cheering jingle when a check matches the expected
// HTTP status code. The real implementation needs an audio backend
// (CGO on Linux/macOS, pure Go on Windows); builds without one get a
// silent no-op player so cross-compiled binaries still work.
type player interface {
	// Play plays the jingle and blocks until it finishes.
	Play()
}
