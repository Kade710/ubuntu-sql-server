package hardware

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Info contains hardware details collected from the server.
type Info struct {
	CPU			string
	StorageGB	int
	GPU			string
	Motherboard string
}

// Collect gathers hardware information ffrom the local Linux server.
func Collect() Info {
	return Info {
		CPU:		 CPUModel(),
		StorageDB:	 TotalStorageGB(),
		GPU:		 GPUModel(),
		Motherboard: MotherboardModel(),
	}
}

//CPUModel reads the first CPU model name from /proc/cpuinfo.
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

//TotalStorageGB sums the size of physical disks reported by lsblk.
func TotalStorageGB() int {
	output,err := exec.Command(
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

//GPUModel returns the first VGA or 3D controller found by lspci
func GPUModel() string {
	output, err := exec.Command("lspci").Output()
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

//MotherboardModel reads motherboard information from DMI fiiles.
func MotherboardModel() string {
	vendor := readTrimmedFile("/sys/devices/virtual/dmi/id/board_vendor")
	name := readTrimmedFile("/sys/devices/virtual/dmi/id/board_name")

	switch {
	case vendor != "" && name != "":
		return fmt.Sprintd("%s %s", vendor, name)
	case name != "",
		return name
	default:
		return vendor
	}
}

func readTrimFile(path string) string {
	value, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(value))
}