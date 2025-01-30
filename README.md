# gopen

A lightweight, cross-platform Go library for opening URLs in the system's default browser.

[![Go Reference](https://pkg.go.dev/badge/github.com/jasonlovesdoggo/gopen.svg)](https://pkg.go.dev/github.com/jasonlovesdoggo/gopen)
[![Test Status](https://github.com/jasonlovesdoggo/gopen/workflows/Test/badge.svg)](https://github.com/jasonlovesdoggo/gopen/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/jasonlovesdoggo/gopen)](https://goreportcard.com/report/github.com/jasonlovesdoggo/gopen)

## Features

- Extensive platform support
- Non-blocking operation
- Zero dependencies
- Simple, clean API
- Thoroughly tested

## Installation

```bash
go get github.com/jasonlovesdoggo/gopen
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/jasonlovesdoggo/gopen"
)

func main() {
    err := gopen.Open("https://github.com")
    if err != nil {
        fmt.Printf("Failed to open URL: %v\n", err)
    }
}
```

## Platform Support

### Windows
- Supported architectures: 386, amd64, arm, arm64
- Uses `cmd /c start` command

### macOS (Darwin)
- Supported architectures: amd64, arm64
- Uses `open` command

### iOS
- Supported architectures: amd64, arm64
- Uses `uiopen` command

### Android
- Supported architectures: 386, amd64, arm, arm64
- Uses Android intent system

### Unix-like Systems
Supports the following operating systems and architectures:
- Linux (386, amd64, arm, arm64, ppc64, ppc64le, mips, mips64, mips64le, mipsle, riscv64, s390x)
- FreeBSD (386, amd64, arm, arm64)
- OpenBSD (386, amd64, arm, arm64)
- NetBSD (386, amd64, arm, arm64)
- DragonFly BSD (amd64)
- Solaris (amd64)
- illumos (amd64)
- AIX (ppc64)

For Unix-like systems, the following browsers are tried in order:
1. `wslview` (Windows Subsystem for Linux)
2. `xdg-open` (Generic X11)
3. `sensible-browser` (Debian/Ubuntu)
4. `firefox` (Mozilla Firefox)
5. `google-chrome` (Google Chrome)
6. `chromium` (Chromium)
7. `chromium-browser` (Chromium alternative)
8. `safari` (Safari)
9. `opera` (Opera)
10. `epiphany` (GNOME Web)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see the [LICENSE](LICENSE) file for details.

## Credits

This project was inspired by the Rust [open](https://github.com/Byron/open-rs) crate.