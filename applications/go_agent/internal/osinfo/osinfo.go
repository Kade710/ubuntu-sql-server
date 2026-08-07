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

	var distribution string
	var version string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "NAME=") {
			distribution = strings.Trim(
				strings.TrimPrefix(line, "NAME="),
				`"`,
			)
		}

		if strings.HasPrefix(line, "VERSION=") {
			version = strings.Trim(
				strings.TrimPrefix(line, "VERSION="),
				`"`,
			)
		}
	}

	return distribution, version
}

func kernelVersion() string {
	var uname syscall.Utsname

	if err := syscall.Uname(&uname); err != nil {
		return ""
	}

	return "Linux " + charsToString(uname.Release[:])
}

func charsToString(chars []int8) string {
	result := make([]byte, 0, len(chars))

	for _, char := range chars {
		if char == 0 {
			break
		}

		result = append(result, byte(char))
	}

	return string(result)
}

func architecture() string {
	switch runtime.GOARCH {
	case "amd64":
		return "x86_64"
	case "arm64":
		return "aarch64"
	default:
		return runtime.GOARCH
	}
}
