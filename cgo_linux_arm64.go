//go:build linux && arm64

package bindings

// Linux ARM64: Link against bundled static libraries

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/linux_arm64 -Wl,--whole-archive -lfaiss_c -Wl,--no-whole-archive -lfaiss
#cgo LDFLAGS: -lstdc++ -lm -lpthread -ldl
#cgo LDFLAGS: -lgomp -lgfortran
*/
import "C"
