package system

import (
	"runtime"
	"testing"
)

func TestOperatingSystem(t *testing.T) {
	got := OperatingSystem()
	want := runtime.GOOS

	if got != want {
		t.Fatalf("OperatingSystem() = %q, want %q", got, want)
	}
}

func TestArchitecture(t *testing.T) {
	got := Architecture()
	want := runtime.GOARCH

	if got != want {
		t.Fatalf("Architecture() = %q, want %q", got, want)
	}
}
