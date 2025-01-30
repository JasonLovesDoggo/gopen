package gopen

import (
	"errors"
	"regexp"
	"runtime"
)

var (
	// ErrUnsupportedOS is returned when the operating system is not supported
	ErrUnsupportedOS = errors.New("unsupported operating system")
	// ErrOpenFailed is returned when all attempts to open the URL fail
	ErrOpenFailed = errors.New("failed to open URL: no available browser command succeeded")
	ErrInvalidURL = errors.New("invalid URL")
)

// Open attempts to open the provided URL in the system's default browser.
// The function is non-blocking and returns immediately after launching the browser.
//
// On Windows, it uses the 'start' command
// On macOS, it uses the 'open' command
// On iOS, it uses the 'uiopen' command
// On Linux/Unix, it tries several common browsers
//
// Returns an error if the URL cannot be opened or if the OS is not supported.
func Open(url string) error {
	if url == "" {
		return errors.New("empty URL")
	}

	// Simple regex to check if the URL is valid src: https://regexr.com/39p0t
	urlRegex, err := regexp.Compile(`(https?:\/\/)?([\w\-])+\.{1}([a-zA-Z]{2,63})([\/\w-]*)*\/?\??([^#\n\r]*)?#?([^\n\r]*)`)
	if err != nil {
		panic(err) // This should never happen
	}

	if !urlRegex.MatchString(url) {
		return ErrInvalidURL
	}

	switch runtime.GOOS {
	case "darwin": // darwin/amd64, darwin/arm64
		return Darwin(url)
	case "windows": // windows/386, windows/amd64, windows/arm, windows/arm64
		return Windows(url)
	case "ios": // ios/amd64, ios/arm64
		return IOS(url)
	case "android": // android/386, android/amd64, android/arm, android/arm64
		return Android(url)
	case "linux", // linux/386, linux/amd64, linux/arm, linux/arm64, linux/ppc64, linux/ppc64le, linux/mips, linux/mips64, linux/mips64le, linux/mipsle, linux/riscv64, linux/s390x
		"freebsd",   // freebsd/386, freebsd/amd64, freebsd/arm, freebsd/arm64
		"openbsd",   // openbsd/386, openbsd/amd64, openbsd/arm, openbsd/arm64
		"netbsd",    // netbsd/386, netbsd/amd64, netbsd/arm, netbsd/arm64
		"dragonfly", // dragonfly/amd64
		"solaris",   // solaris/amd64
		"illumos",   // illumos/amd64
		"aix":       // aix/ppc64
		return Unix(url)

	default:
		return ErrUnsupportedOS
	}
}
