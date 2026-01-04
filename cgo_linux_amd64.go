//go:build linux && amd64

package bindings

// Linux AMD64: Link against bundled static libraries
// The lib/linux_amd64/ directory should contain:
// - libfaiss.a (merged with OpenBLAS)
// - libfaiss_c.a
// - libgfortran.a (optional, for full static)
// - libgomp.a (optional, for full static)
//
// If libgfortran.a and libgomp.a are bundled, we use them.
// Otherwise, fall back to system libraries.

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/linux_amd64 -Wl,--whole-archive -lfaiss_c -Wl,--no-whole-archive -lfaiss
#cgo LDFLAGS: -lstdc++ -lm -lpthread -ldl
#cgo LDFLAGS: -lgomp -lgfortran
*/
import "C"
