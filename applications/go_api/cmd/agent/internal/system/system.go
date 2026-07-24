package system

import (
    "os"
    "runtime"
)

func  Hostname() (string, error) {
    return os.Hostname()
}

func OperatingSystem() string {
    return runtime.GOOS
}

func Architecture() string {
    return runtime.GOARCH
}