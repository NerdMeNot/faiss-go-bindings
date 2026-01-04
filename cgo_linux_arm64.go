//go:build linux && arm64

package bindings

// Linux ARM64: Link against bundled static libraries
// The lib/linux_arm64/ directory should contain:
// - libfaiss.a (merged with OpenBLAS)
// - libfaiss_c.a
// - libfaiss_go_ext.a (custom Go extensions)
//
// Runtime dependencies: libgomp (OpenMP), libgfortran

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/linux_arm64 -lfaiss_c -lfaiss_go_ext -lfaiss
#cgo LDFLAGS: -lstdc++ -lm -lpthread -ldl
#cgo LDFLAGS: -lgomp -lgfortran
*/
import "C"
