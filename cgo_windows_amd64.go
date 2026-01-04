//go:build windows && amd64

package bindings

// Windows AMD64: Link against bundled static libraries
// The lib/windows_amd64/ directory should contain:
// - faiss.lib
// - faiss_c.lib
// - openblas.lib (bundled from vcpkg)
//
// Note: Windows support is experimental. Building requires MinGW-w64.
// libfaiss_go_ext is not yet supported on Windows.

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/windows_amd64 -lfaiss_c -lfaiss -lopenblas
#cgo LDFLAGS: -lstdc++ -lm -lpthread
*/
import "C"
