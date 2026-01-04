//go:build linux && amd64

package bindings

// Linux AMD64: Link against bundled static libraries
// The lib/linux_amd64/ directory should contain:
// - libfaiss.a (merged with OpenBLAS)
// - libfaiss_c.a
// - libfaiss_go_ext.a (custom Go extensions)
//
// Runtime dependencies: libgomp (OpenMP), libgfortran

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/linux_amd64 -Wl,--whole-archive -lfaiss_c -lfaiss_go_ext -Wl,--no-whole-archive -lfaiss
#cgo LDFLAGS: -lstdc++ -lm -lpthread -ldl
#cgo LDFLAGS: -lgomp -lgfortran
*/
import "C"
