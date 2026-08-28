package inventory

import (
	"bufio"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/hardware"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/system"
)

// Server contains the inventory collected by the Go agent.
type Server struct {
	Hostname        string
	IPAddress       string
	OperatingSystem string
	RAMGB           int

	CPU         string
	StorageGB   int
	GPU         string
	Motherboard string
}

// Collect gathers basic inventory information from the server.
func Collect() (Server, error) {
	hostname, err := system.Hostname()
	if err != nil {
		return Server{}, err
	}

	memoryGB, err := system.TotalMemoryGB()
	if err != nil {
		return Server{}, err
	}

	hw := hardware.Collect()

	return Server{
		Hostname:        strings.TrimSpace(hostname),
		IPAddress:       primaryIPv4(),
		OperatingSystem: prettyOperatingSystem(),
		RAMGB:           int(memoryGB + 0.5),

		CPU:         hw.CPU,
		StorageGB:   hw.StorageGB,
		GPU:         hw.GPU,
		Motherboard: hw.Motherboard,
	}, nil
}

func prettyOperatingSystem() string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return runtime.GOOS
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "PRETTY_NAME=") {
			value := strings.TrimPrefix(line, "PRETTY_NAME=")
			value = strings.TrimSpace(value)
			value = strings.Trim(value, `"`)

			if value != "" {
				return value
			}
		}
	}

	return runtime.GOOS
}

func primaryIPv4() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	var privateFallback string
	var publicFallback string

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 ||
			iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addresses, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, address := range addresses {
			var ip net.IP

			switch value := address.(type) {
			case *net.IPNet:
				ip = value.IP
			case *net.IPAddr:
				ip = value.IP
			default:
				continue
			}

			ip = ip.To4()
			if ip == nil ||
				ip.IsLoopback() ||
				ip.IsUnspecified() ||
				ip.IsMulticast() {
				continue
			}

			ipString := ip.String()

			if ip.IsPrivate() {
				if !isVirtualInterface(iface.Name) {
					return ipString
				}

				if privateFallback == "" {
					privateFallback = ipString
				}

				continue
			}

			if publicFallback == "" && !isVirtualInterface(iface.Name) {
				publicFallback = ipString
			}
		}
	}

	if privateFallback != "" {
		return privateFallback
	}

	return publicFallback
}

func isVirtualInterface(name string) bool {
	name = strings.ToLower(strings.TrimSpace(name))

	virtualPrefixes := []string{
		"docker",
		"br-",
		"veth",
		"virbr",
		"tailscale",
		"tun",
		"tap",
	}

	for _, prefix := range virtualPrefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}

	return false
}
