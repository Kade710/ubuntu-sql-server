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

// Components convert collected hardware information into database records.
func Components(info Info, ramGB int) []Component {
	return []Component{
		cpuComponent(info.CPU),
		{
			Type:          "RAM",
			Specification: fmt.Sprintf("%d GB", ramGB),
		},
		motherboardComponent(info.Motherboard),
		gpuComponent(info.GPU),
	}
}

func cpuComponent(cpu string) Component {
	component := Component{
		Type:          "CPU",
		Model:         cpu,
		Specification: cpu,
	}

	lower := strings.ToLower(cpu)

	switch {
	case strings.Contains(lower, "intel"):
		component.Manufacturer = "Intel"
	case strings.Contains(lower, "amd"):
		component.Manufacturer = "AMD"
	}

	return component
}

func gpuComponent(gpu string) Component {
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
	case strings.Contains(lower, "amd"):
		component.Manufacturer = "AMD"
	}

	return component
}

func motherboardComponent(board string) Component {
	component := Component{
		Type:          "Motherboard",
		Model:         board,
		Specification: board,
	}

	lower := strings.ToLower(board)

	switch {
	case strings.HasPrefix(lower, "msi"):
		component.Manufacturer = "MSI"
	case strings.HasPrefix(lower, "asus"):
		component.Manufacturer = "ASUS"
	case strings.HasPrefix(lower, "gigabyte"):
		component.Manufacturer = "Gigabyte"
	case strings.HasPrefix(lower, "asrock"):
		component.Manufacturer = "ASRock"
	}

	return component
}
