//go:build linux && arm64

package bindings

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/linux_arm64 -lfaiss_c -lfaiss -lopenblas -lgfortran -lgomp -lpthread
*/
import "C"
