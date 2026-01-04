# Runtime Dependencies

This document explains the runtime dependencies required for faiss-go-bindings.

## Summary

| Platform | Required Runtime Packages |
|----------|--------------------------|
| **Linux AMD64/ARM64** | `libgomp1` (OpenMP runtime) |
| **macOS Intel** | `libomp` from Homebrew |
| **macOS Apple Silicon** | `libomp` from Homebrew |

## Why Runtime Dependencies?

FAISS uses:
1. **BLAS/LAPACK** - For linear algebra operations
2. **OpenMP** - For multi-threaded parallel operations

### What's Bundled (Static)

The pre-built libraries include:
- `libfaiss.a` - Core FAISS with OpenBLAS statically linked
- `libfaiss_c.a` - C API wrapper

### What's NOT Bundled (Dynamic)

- **OpenMP runtime** (`libgomp` on Linux, `libomp` on macOS)
  - Static linking of OpenMP is problematic and can cause issues
  - The dynamic library is small (~200KB) and widely available

## Installation

### Linux (Ubuntu/Debian)
```bash
sudo apt-get install libgomp1
```

### Linux (RHEL/CentOS/Fedora)
```bash
sudo dnf install libgomp
```

### macOS
```bash
brew install libomp
```

## Verification

To verify dependencies are available:

### Linux
```bash
ldconfig -p | grep gomp
# Should show: libgomp.so.1 ...
```

### macOS
```bash
ls /opt/homebrew/opt/libomp/lib/libomp.dylib  # Apple Silicon
ls /usr/local/opt/libomp/lib/libomp.dylib     # Intel
```

## Building Without OpenMP (Advanced)

If you want truly zero runtime dependencies, you can rebuild FAISS without OpenMP:

1. Set `FAISS_OPT_LEVEL=generic` and disable OpenMP in CMake
2. This trades parallelism for zero dependencies
3. Performance will be significantly reduced for large operations

This is NOT recommended for production use.

## Future Improvements

We're exploring:
1. Bundling static OpenMP (requires careful testing)
2. Optional builds without OpenMP for embedded use
3. Container images with all dependencies pre-installed
