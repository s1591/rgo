//go:build (!cgo || nocgo) && unix && !darwin

package mpv

import (
	"fmt"

	"github.com/ebitengine/purego"
)

const (
	libname = "libmpv.so"
)

// loadLibrary loads the so and panics on error.
func loadLibrary() uintptr {
	handle, err := purego.Dlopen(libname, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(fmt.Errorf("cannot load library: %w", err))
	}

	return uintptr(handle)
}
