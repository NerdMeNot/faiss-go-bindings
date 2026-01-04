# CLAUDE.md

This file provides guidance to Claude Code when working with this repository.

## Overview

faiss-go-bindings provides pre-built FAISS static libraries for Go. This is a **bindings-only** module - it contains no Go wrapper code, just CGO configuration and static libraries.

## Project Structure

```
faiss-go-bindings/
├── bindings.go           # Main CGO declarations and C API types
├── bindings_test.go      # Smoke tests for library linking
├── cgo_darwin_amd64.go   # macOS Intel CGO flags
├── cgo_darwin_arm64.go   # macOS Apple Silicon CGO flags
├── cgo_linux_amd64.go    # Linux AMD64 CGO flags
├── cgo_linux_arm64.go    # Linux ARM64 CGO flags
├── lib/                  # Pre-built static libraries
│   ├── darwin_amd64/
│   ├── darwin_arm64/
│   ├── linux_amd64/
│   └── linux_arm64/
├── include/              # C header files
├── c_api_ext/            # Custom C extensions source
└── scripts/              # Build scripts for static libraries
```

## Supported Platforms

- linux/amd64, linux/arm64
- darwin/amd64, darwin/arm64

## Key Commands

```bash
# Build
go build -v ./...

# Test (smoke tests only)
go test -v ./...

# Test with coverage
go test -coverprofile=coverage.out ./...
```

## How It Works

1. Platform-specific `cgo_*.go` files set CGO LDFLAGS
2. Static libraries in `lib/<platform>/` are linked at build time
3. The `bindings.go` file declares C API types and test helpers

## CGO Configuration

Each platform file sets:
- `CFLAGS`: Include path to headers
- `LDFLAGS`: Library paths and link flags

Example (linux_amd64):
```go
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/linux_amd64 -lfaiss_c -lfaiss_go_ext -lfaiss
#cgo LDFLAGS: -lstdc++ -lm -lpthread -ldl -lgomp -lgfortran
```

## Static Libraries

Each platform needs:
- `libfaiss.a` - Core FAISS (with OpenBLAS merged on Linux)
- `libfaiss_c.a` - FAISS C API
- `libfaiss_go_ext.a` - Custom Go extensions

## Workflows

- `build-static-libs.yml` - Rebuild static libraries for all platforms
- `release.yml` - Create GitHub releases
- `check-faiss-releases.yml` - Monitor for new FAISS versions
- `ci.yml` - Run tests on push

## Important Notes

- This module is meant to be imported as a dependency by faiss-go
- Do NOT add Go wrapper code here - that belongs in faiss-go
- Static libraries must be rebuilt when FAISS version changes
- Linux builds merge OpenBLAS statically; macOS uses Accelerate framework

## Runtime Dependencies

- **Linux**: `libgomp1` (OpenMP runtime)
- **macOS**: `libomp` from Homebrew
