# faiss-go-bindings

Pre-built FAISS static libraries for Go.

This module provides pre-built static libraries for [FAISS](https://github.com/facebookresearch/faiss) (Facebook AI Similarity Search) that can be used with [faiss-go](https://github.com/NerdMeNot/faiss-go).

## Usage

This module is designed to be imported with a blank identifier in your Go code:

```go
import _ "github.com/NerdMeNot/faiss-go-bindings"
```

The CGO directives in this package configure the linker to use the pre-built FAISS libraries bundled with this module.

## Supported Platforms

| Platform | Architecture | Status |
|----------|--------------|--------|
| Linux    | amd64        | Supported |
| Linux    | arm64        | Supported |
| macOS    | amd64        | Supported |
| macOS    | arm64        | Supported |
| Windows  | amd64        | Supported |

## Runtime Dependencies

The FAISS core and OpenBLAS are statically linked. Only the OpenMP runtime is required:

| Platform | Install Command |
|----------|-----------------|
| **Linux** | `sudo apt-get install libgomp1` |
| **macOS** | `brew install libomp` |

See [DEPENDENCIES.md](DEPENDENCIES.md) for detailed information.

## Building Libraries

To rebuild the static libraries, run the build workflow:

```bash
# Build for current platform
./scripts/build-static-libs.sh

# Or use GitHub Actions workflow
# See .github/workflows/build-static-libs.yml
```

## Module Structure

```
faiss-go-bindings/
├── bindings.go           # CGO declarations and C API types
├── cgo_darwin_amd64.go   # macOS Intel LDFLAGS
├── cgo_darwin_arm64.go   # macOS Apple Silicon LDFLAGS
├── cgo_linux_amd64.go    # Linux AMD64 LDFLAGS
├── cgo_linux_arm64.go    # Linux ARM64 LDFLAGS
├── cgo_windows_amd64.go  # Windows AMD64 LDFLAGS
├── include/              # FAISS C API headers
├── lib/                  # Pre-built static libraries
│   ├── darwin_amd64/
│   ├── darwin_arm64/
│   ├── linux_amd64/
│   ├── linux_arm64/
│   └── windows_amd64/
├── c_api_ext/            # Custom C extensions source
└── scripts/              # Build scripts
```

## Version Information

- FAISS Version: 1.8.0
- Go Version: 1.21+

## License

MIT License - see [LICENSE](LICENSE)

## Related Projects

- [faiss-go](https://github.com/NerdMeNot/faiss-go) - Go bindings for FAISS
- [FAISS](https://github.com/facebookresearch/faiss) - Facebook AI Similarity Search
