package hardware

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const commandTimeout = 5 * time.Second

// Info contains hardware details collected from the server.
type Info struct {
	CPU         string
	StorageGB   int
	GPU         string
	Motherboard string

	PSUManufacturer string
	PSUModel        string
	PSUWatts        int
}

// Collect gathers hardware information from the local Linux server.
func Collect() Info {
	return Info{
		CPU:         CPUModel(),
		StorageGB:   TotalStorageGB(),
		GPU:         GPUModel(),
		Motherboard: MotherboardModel(),

		PSUManufacturer: strings.TrimSpace(os.Getenv("PSU_MANUFACTURER")),
		PSUModel:        strings.TrimSpace(os.Getenv("PSU_MODEL")),
		PSUWatts:        PSUWatts(),
	}
}

// CPUModel reads the first CPU model name from /proc/cpuinfo.
func CPUModel() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "model name") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	return ""
}

// TotalStorageGB sums the size of physical disks reported by lsblk.
func TotalStorageGB() int {
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	output, err := exec.CommandContext(
		ctx,
		"lsblk",
		"-b",
		"-d",
		"-n",
		"-o",
		"SIZE,TYPE",
	).Output()
	if err != nil {
		return 0
	}

	var totalBytes uint64

	scanner := bufio.NewScanner(strings.NewReader(string(output)))

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 || fields[1] != "disk" {
			continue
		}

		size, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			continue
		}

		totalBytes += size
	}

	const bytesPerGB = 1_000_000_000

	return int((totalBytes + bytesPerGB/2) / bytesPerGB)
}

// GPUModel returns the first VGA or 3D controller found by lspci.
func GPUModel() string {
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	output, err := exec.CommandContext(ctx, "lspci").Output()
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "VGA compatible controller") ||
			strings.Contains(line, "3D controller") {

			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}

			return strings.TrimSpace(line)
		}
	}

	return ""
}

// MotherboardModel reads motherboard information from Linux DMI files.
func MotherboardModel() string {
	vendor := readTrimmedFile("/sys/devices/virtual/dmi/id/board_vendor")
	name := readTrimmedFile("/sys/devices/virtual/dmi/id/board_name")

	switch {
	case vendor != "" && name != "":
		return fmt.Sprintf("%s %s", vendor, name)
	case name != "":
		return name
	default:
		return vendor
	}
}

// PSUWatts reads the manually configured PSU wattage.
func PSUWatts() int {
	value := strings.TrimSpace(os.Getenv("PSU_WATTS"))
	if value == "" {
		return 0
	}

	watts, err := strconv.Atoi(value)
	if err != nil || watts < 0 {
		return 0
	}

	return watts
}

func readTrimmedFile(path string) string {
	value, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(value))
}
