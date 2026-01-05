//go:build darwin && amd64

package bindings

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/darwin_amd64 -lfaiss_c -lfaiss -lfaiss_go_ext -L/usr/local/opt/libomp/lib -lomp -lstdc++ -lm -Wl,-framework,Accelerate
*/
import "C"
