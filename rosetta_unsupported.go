// +build !darwin

package rosetta

import (
	"runtime"
)

// Enabled returns true if running in a Rosetta Translated Binary, false otherwise.
func Enabled() bool {
	return false
}

// NativeArch returns the native architecture, even if binary architecture
// is emulated by Rosetta.
func NativeArch() string {
	return runtime.GOARCH
}
