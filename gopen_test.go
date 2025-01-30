package gopen

import (
	"errors"
	"os"
	"runtime"
	"testing"
)

func TestOpen(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "valid http URL",
			url:     "http://example.com",
			wantErr: false,
		},
		{
			name:    "valid https URL",
			url:     "https://example.com",
			wantErr: false,
		},
		{
			name:    "empty URL",
			url:     "",
			wantErr: true,
		},
	}

	// Skip actual browser opening tests if running in CI
	if os.Getenv("CI") != "" {
		t.Skip("Skipping browser tests in CI environment")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Open(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnsupportedOS(t *testing.T) {
	// Only run this test if we're on an unsupported OS
	if runtime.GOOS == "darwin" || runtime.GOOS == "windows" || runtime.GOOS == "linux" ||
		runtime.GOOS == "freebsd" || runtime.GOOS == "netbsd" || runtime.GOOS == "openbsd" ||
		runtime.GOOS == "ios" {
		t.Skip("Skipping unsupported OS test on supported platform")
	}

	err := Open("https://example.com")
	if !errors.Is(err, ErrUnsupportedOS) {
		t.Errorf("Expected ErrUnsupportedOS, got %v", err)
	}
}
