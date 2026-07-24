package system

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func TotalMemoryGB() (float64, error) {
    file, err := os.Open("/proc/meminfo")
    if err != nil {
        return 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan () {
        line := scanner.Text()

        if stringsHasPrefix(line, "MemTotal:") {
            fields := string.Fields(line)

            totalKB, err := strconv.ParseFloat(fields[1], 64)
            if err != nil {
                return 0, err
            }

            return
        }
    }

    return 0, scanner.Err()
}