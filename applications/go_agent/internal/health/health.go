package health

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"syscall"
)

const (
	warningLoad  = 2.0
	criticalLoad = 4.0

	warningMemory  = 60.0
	criticalMemory = 75.0

	warningDisk  = 70.0
	criticalDisk = 80.0
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

	if math.IsNaN(load) || math.IsInf(load, 0) || load < 0 {
		return 0, fmt.Errorf("invalid load average value")
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
	var foundTotal bool
	var foundAvailable bool

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "MemTotal:":
			value, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return 0, fmt.Errorf("parse total memory: %w", err)
			}

			total = value
			foundTotal = true

		case "MemAvailable:":
			value, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return 0, fmt.Errorf("parse available memory: %w", err)
			}

			available = value
			foundAvailable = true
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scan memory info: %w", err)
	}

	if !foundTotal || total == 0 {
		return 0, fmt.Errorf("unable to determine total memory")
	}

	if !foundAvailable {
		return 0, fmt.Errorf("unable to determine available memory")
	}

	if available > total {
		return 0, fmt.Errorf("available memory exceeds total memory")
	}

	used := total - available
	percent := float64(used) / float64(total) * 100

	if percent < 0 || percent > 100 {
		return 0, fmt.Errorf("invalid memory usage percentage")
	}

	return percent, nil
}

func diskUsagePercent(path string) (float64, error) {
	var stat syscall.Statfs_t

	if err := syscall.Statfs(path, &stat); err != nil {
		return 0, fmt.Errorf("read disk usage for %s: %w", path, err)
	}

	total := stat.Blocks
	available := stat.Bavail

	if total == 0 {
		return 0, fmt.Errorf("unable to determine disk size")
	}

	if available > total {
		return 0, fmt.Errorf("available disk blocks exceed total blocks")
	}

	used := total - available
	percent := float64(used) / float64(total) * 100

	if percent < 0 || percent > 100 {
		return 0, fmt.Errorf("invalid disk usage percentage")
	}

	return percent, nil
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

	if math.IsNaN(seconds) || math.IsInf(seconds, 0) || seconds < 0 {
		return 0, fmt.Errorf("invalid uptime value")
	}

	return seconds / 3600, nil
}

func overallStatus(load, memoryPercent, diskPercent float64) string {
	if load >= criticalLoad ||
		memoryPercent >= criticalMemory ||
		diskPercent >= criticalDisk {
		return "CRITICAL"
	}

	if load >= warningLoad ||
		memoryPercent >= warningMemory ||
		diskPercent >= warningDisk {
		return "WARNING"
	}

	return "HEALTHY"
}
