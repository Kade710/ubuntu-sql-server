package health

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

// Status contains current server health information.
type Status struct {
	Load1         float64
	MemoryPercent float64
	DiskPercent   float64
	UptimeHours   float64
	Overall       string
}

// Collect gathers basic Linux health metrics.
func Collect() (Status, error) {
	load, err := loadAverage()
	if err != nil {
		return Status{}, err
	}

	memory, err := memoryUsagePercent()
	if err != nil {
		return Status{}, err
	}

	disk, err := diskUsagePercent("/")
	if err != nil {
		return Status{}, err
	}

	uptime, err := uptimeHours()
	if err != nil {
		return Status{}, err
	}

	return Status{
		Load1:         load,
		MemoryPercent: memory,
		DiskPercent:   disk,
		UptimeHours:   uptime,
		Overall:       overallStatus(load, memory, disk),
	}, nil
}

func loadAverage() (float64, error) {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return 0, fmt.Errorf("read load average: %w", err)
	}

	fields := strings.Fields(string(data))
	if len(fields) == 0 {
		return 0, fmt.Errorf("invalid /proc/loadavg")
	}

	load, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, fmt.Errorf("parse load average: %w", err)
	}

	return load, nil
}

func memoryUsagePercent() (float64, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return 0, fmt.Errorf("read memory info: %w", err)
	}
	defer file.Close()

	var total uint64
	var available uint64

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "MemTotal:":
			total, _ = strconv.ParseUint(fields[1], 10, 64)
		case "MemAvailable:":
			available, _ = strconv.ParseUint(fields[1], 10, 64)
		}
	}

	if total == 0 {
		return 0, fmt.Errorf("unable to determine total memory")
	}

	used := total - available

	return float64(used) / float64(total) * 100, nil
}

func diskUsagePercent(path string) (float64, error) {
	var stat syscall.Statfs_t

	if err := syscall.Statfs(path, &stat); err != nil {
		return 0, fmt.Errorf("read disk usage: %w", err)
	}

	total := stat.Blocks
	available := stat.Bavail

	if total == 0 {
		return 0, fmt.Errorf("unable to determine disk size")
	}

	used := total - available

	return float64(used) / float64(total) * 100, nil
}

func uptimeHours() (float64, error) {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return 0, fmt.Errorf("read uptime: %w", err)
	}

	fields := strings.Fields(string(data))
	if len(fields) == 0 {
		return 0, fmt.Errorf("invalid /proc/uptime")
	}

	seconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, fmt.Errorf("parse uptime: %w", err)
	}

	return seconds / 3600, nil
}

func overallStatus(memoryPercent, diskPercent float64) string {
	if load >= 8.0 || memoryPercent >= 90 || diskPercent >= 90 {
		return "CRITICAL"
	}

	if load >= 4.0 || memoryPercent >= 75 || diskPercent >= 80 {
		return "WARNING"
	}

	return "HEALTHY"
}
