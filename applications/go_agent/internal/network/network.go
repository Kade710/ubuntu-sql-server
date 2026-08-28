package network

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Interface contains network information collected from Linux.
type Interface struct {
	Name        string
	MACAddress  string
	IPAddress   string
	NetworkType string
	SpeedMbps   int
}

// Collect gathers active network interfaces with IPv4 addresses.
func Collect() ([]Interface, error) {
	systemInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("list network interfaces: %w", err)
	}

	var collected []Interface

	for _, iface := range systemInterfaces {
		if iface.Flags&net.FlagUp == 0 ||
			iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		ipAddress := ipv4Address(iface)
		if ipAddress == "" {
			continue
		}

		collected = append(collected, Interface{
			Name:        strings.TrimSpace(iface.Name),
			MACAddress:  iface.HardwareAddr.String(),
			IPAddress:   ipAddress,
			NetworkType: interfaceType(iface.Name),
			SpeedMbps:   interfaceSpeed(iface.Name),
		})
	}

	return collected, nil
}

func ipv4Address(iface net.Interface) string {
	addresses, err := iface.Addrs()
	if err != nil {
		return ""
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

		return ip.String()
	}

	return ""
}

func interfaceType(name string) string {
	name = strings.TrimSpace(name)
	lowerName := strings.ToLower(name)

	virtualPrefixes := []string{
		"docker",
		"br-",
		"veth",
		"virbr",
	}

	for _, prefix := range virtualPrefixes {
		if strings.HasPrefix(lowerName, prefix) {
			return "Virtual"
		}
	}

	tunnelPrefixes := []string{
		"tailscale",
		"tun",
		"tap",
	}

	for _, prefix := range tunnelPrefixes {
		if strings.HasPrefix(lowerName, prefix) {
			return "Tunnel"
		}
	}

	wirelessPath := filepath.Join("/sys/class/net", name, "wireless")

	if _, err := os.Stat(wirelessPath); err == nil {
		return "Wireless"
	}

	return "Ethernet"
}

func interfaceSpeed(name string) int {
	name = strings.TrimSpace(name)
	if name == "" {
		return 0
	}

	speedPath := filepath.Join("/sys/class/net", name, "speed")

	value, err := os.ReadFile(speedPath)
	if err != nil {
		return 0
	}

	speed, err := strconv.Atoi(strings.TrimSpace(string(value)))
	if err != nil || speed < 0 {
		return 0
	}

	return speed
}
