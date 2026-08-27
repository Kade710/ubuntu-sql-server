package hardware

import (
	"fmt"
	"strings"
)

// Component represents a hardware component stored in PostgreSQL.
type Component struct {
	Type          string
	Manufacturer  string
	Model         string
	Specification string
}

// Components converts collected hardware information into database records.
func Components(info Info, ramGB int) []Component {
	if ramGB < 0 {
		ramGB = 0
	}

	return []Component{
		cpuComponent(info.CPU),
		{
			Type:          "RAM",
			Specification: fmt.Sprintf("%d GB", ramGB),
		},
		motherboardComponent(info.Motherboard),
		gpuComponent(info.GPU),
		psuComponent(info),
	}
}

func cpuComponent(cpu string) Component {
	cpu = strings.TrimSpace(cpu)

	component := Component{
		Type:          "CPU",
		Model:         cpu,
		Specification: cpu,
	}

	lower := strings.ToLower(cpu)

	switch {
	case strings.Contains(lower, "intel"):
		component.Manufacturer = "Intel"
	case strings.Contains(lower, "amd"),
		strings.Contains(lower, "advanced micro devices"):
		component.Manufacturer = "AMD"
	}

	return component
}

func gpuComponent(gpu string) Component {
	gpu = strings.TrimSpace(gpu)

	component := Component{
		Type:          "GPU",
		Model:         gpu,
		Specification: gpu,
	}

	lower := strings.ToLower(gpu)

	switch {
	case strings.Contains(lower, "nvidia"):
		component.Manufacturer = "NVIDIA"
	case strings.Contains(lower, "intel"):
		component.Manufacturer = "Intel"
	case strings.Contains(lower, "amd"),
		strings.Contains(lower, "advanced micro devices"):
		component.Manufacturer = "AMD"
	}

	return component
}

func psuComponent(info Info) Component {
	component := Component{
		Type:			"PSU",
		Manufacturer:	strings.TrimSpace(info.PSUManufacturer),
		Model:			strings.TrimSpace(info.PSUModel),
	}

	if info.PSUWatts > 0 {
		component.Specification = fmt.Sprintf("%d W", info.PSUWatts)
	}

	return component
}

func motherboardComponent(board string) Component {
	board = strings.TrimSpace(board)

	component := Component{
		Type:          "Motherboard",
		Model:         board,
		Specification: board,
	}

	lower := strings.ToLower(board)

	switch {
	case strings.HasPrefix(lower, "msi"),
		strings.Contains(lower, "micro-star"):
		component.Manufacturer = "MSI"
	case strings.HasPrefix(lower, "asus"),
		strings.Contains(lower, "asustek"):
		component.Manufacturer = "ASUS"
	case strings.HasPrefix(lower, "gigabyte"):
		component.Manufacturer = "Gigabyte"
	case strings.HasPrefix(lower, "asrock"):
		component.Manufacturer = "ASRock"
	}

	return component
}
