package system

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func Hostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", fmt.Errorf("get hostname: %w", err)
	}

	hostname = strings.TrimSpace(hostname)
	if hostname == "" {
		return "", fmt.Errorf("hostname is empty")
	}

	return hostname, nil
}

func OperatingSystem() string {
	return runtime.GOOS
}

func Architecture() string {
	return runtime.GOARCH
}
