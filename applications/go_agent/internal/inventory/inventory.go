package inventory

import (
	"bufio"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/system"
)

// Server contains the inventory collected by the Go agent.
type Server struct {
	Hostname		string
	IPAdress		string
	OperatingSystem string
	RAMGB			int
}

// Collect gathers basic inventory information from the server.
func Collect() (Server, error) {
	hostnamae, err := system.Hostname()
	if err != nil {
		return Server{}, err
	}

	memoryGB, err := system.TotalMemoryGB()
	if err != nil{
		return Server{}, err
	}

	return Server{
		Hostname:		 hostname,
		IPAddress:		 primaryIPv4(),
		OperatingSystem: prettyOperatingSystem(),
		RAMGB:			 int(memoryGB + 0.5)
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
		line := scanner.Text()

		if strings.HasPrefix(line, "PRETTY_NAME=") {
			value := strings.TrimPrefix(line, "PRETTY_NAME=")
			return strings.Trim(value, `"`)
		}
	}

	return runtime.GOOS
}

func primaryIPv4() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	var fallback string
	for _, iface := range interfaces {
		if iface.Flags&net.Flagup == 0 || iface.Flags&net.FlagLoopback != 0 {
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
			}

			ip = ip.To4()
			if ip == nil || ip.IsLooback() {
				continue
			}

			if ip.IsPrivate() {
				return ip.String()
			}

			if fallback == "" {
				fallback = ip.String()
			}
		}
	}

	return fallback
}