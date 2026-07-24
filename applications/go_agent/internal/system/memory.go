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
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "MemTotal:") {
			fields := strings.Fields(line)

			if len(fields) < 2 {
				return 0, fmt.Errorf("invalid MemTotal entry")
			}

			totalKB, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				return 0, err
			}

			return totalKB / 1024 / 1024, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return 0, fmt.Errorf("MemTotal not found")
}
