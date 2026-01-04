//go:build windows && amd64

package bindings

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/windows_amd64 -lfaiss_c -lfaiss -lopenblas -lgfortran -lquadmath -lpthread
*/
import "C"
