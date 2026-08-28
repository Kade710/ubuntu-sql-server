package system

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TotalMemoryGB() (float64, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return 0, fmt.Errorf("open /proc/meminfo: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if !strings.HasPrefix(line, "MemTotal:") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			return 0, fmt.Errorf("invalid MemTotal entry: %q", line)
		}

		totalKB, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return 0, fmt.Errorf("parse MemTotal value %q: %w", fields[1], err)
		}

		if totalKB <= 0 {
			return 0, fmt.Errorf("invalid MemTotal value: %.0f kB", totalKB)
		}

		return totalKB / 1024 / 1024, nil
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scan /proc/meminfo: %w", err)
	}

	return 0, fmt.Errorf("MemTotal not found in /proc/meminfo")
}
