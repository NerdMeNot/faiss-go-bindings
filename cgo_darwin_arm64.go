//go:build darwin && arm64

package bindings

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/darwin_arm64 -lfaiss_c -lfaiss -lfaiss_go_ext -L/opt/homebrew/opt/libomp/lib -lomp -Wl,-framework,Accelerate
*/
import "C"
