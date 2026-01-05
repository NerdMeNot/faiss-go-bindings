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
#cgo LDFLAGS: -L${SRCDIR}/lib/linux_amd64 -lfaiss_c -lfaiss_go_ext -lfaiss
#cgo LDFLAGS: -lpthread -ldl -lgomp -lgfortran
*/
import "C"
