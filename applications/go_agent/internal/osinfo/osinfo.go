package osinfo

import (
	"bufio"
	"os"
	"runtime"
	"strings"
	"syscall"
)

// Info contains operating system details.
type Info struct {
	Distribution string
	Version      string
	Kernel       string
	Architecture string
}

// Collect gathers operating system details from Linux.
func Collect() Info {
	distribution, version := readOSRelease()

	return Info{
		Distribution: distribution,
		Version:      version,
		Kernel:       kernelVersion(),
		Architecture: architecture(),
	}
}

func readOSRelease() (string, string) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return runtime.GOOS, ""
	}
	defer file.Close()

	distribution := runtime.GOOS
	var version string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}

		value = strings.TrimSpace(value)
		value = strings.Trim(value, `"'`)

		switch key {
		case "NAME":
			if value != "" {
				distribution = value
			}

		case "VERSION":
			version = value
		}
	}

	return distribution, version
}

func kernelVersion() string {
	var uname syscall.Utsname

	if err := syscall.Uname(&uname); err != nil {
		return ""
	}

	release := strings.TrimSpace(charsToString(uname.Release[:]))
	if release == "" {
		return ""
	}

	return "Linux " + release
}

func charsToString(chars []int8) string {
	result := make([]byte, 0, len(chars))

	for _, char := range chars {
		if char == 0 {
			break
		}

		result = append(result, byte(char))
	}

	return strings.TrimSpace(string(result))
}

func architecture() string {
	switch runtime.GOARCH {
	case "amd64":
		return "x86_64"
	case "arm64":
		return "aarch64"
	default:
		return strings.TrimSpace(runtime.GOARCH)
	}
}