package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/config"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/database"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/inventory"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/network"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/system"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("---------------------------")
		fmt.Println(" Ubuntu SQL Server Go Agent")
		fmt.Println("---------------------------")
		fmt.Println("1. Show System Information")
		fmt.Println("2. Show Database Configuration")
		fmt.Println("3. Test Database Connection")
		fmt.Println("4. Register Server Inventory")
		fmt.Println("5. Register Network Interfaces")
		fmt.Println("0. exit")
		fmt.Println("\nSelect an option")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			showSystemInfo()

		case "2":
			showDatabaseConfig()

		case "3":
			testDatabaseConnection()

		case "4":
			registerServerInventory()

		case "5":
			registerNetworkInterfaces()

		case "0":
			fmt.Println("\nSee Ya!")
			return

		default:
			fmt.Println("\nTry Again!")
		}
	}
}

func showSystemInfo() {
	fmt.Println("\nSystem Information")
	fmt.Println("---")

	hostname, err := system.Hostname()
	if err != nil {
		fmt.Println("Hostname:", err)
	} else {
		fmt.Println("Hostname:", hostname)
	}

	fmt.Println("Operating System:", system.OperatingSystem())
	fmt.Println("Architecture:", system.Architecture())

	memoryGB, err := system.TotalMemoryGB()
	if err != nil {
		fmt.Println("Memory:", err)
	} else {
		fmt.Printf("Total Memory: %.2f GB\n", memoryGB)
	}
}

func showDatabaseConfig() {
	fmt.Println("\nDatabase Configuration")
	fmt.Println("---")

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Configuration error:", err)
		return
	}

	fmt.Println("Database host:", cfg.DBHost)
	fmt.Println("Database port:", cfg.DBPort)
	fmt.Println("Database name:", cfg.DBName)
	fmt.Println("Database user:", cfg.DBUser)
	fmt.Println("Database password: configured")
}

func testDatabaseConnection() {
	fmt.Println("\nDatabase Connection Test")
	fmt.Println("---")

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}

	db, err := database.Connect(cfg)
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer db.Close()

	fmt.Println("PostgreSQL connection successful.")
}

func registerServerInventory() {
	fmt.Println("\nRegister Server Inventory")
	fmt.Println("---")

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Configuration error:", err)
		return
	}

	serverInfo, err := inventory.Collect()
	if err != nil {
		fmt.Println("Inventory collection failed:", err)
		return
	}

	db, err := database.Connect(cfg)
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer db.Close()

	serverID, err := database.RegisterServer(db, serverInfo)
	if err != nil {
		fmt.Println("Database update failed:", err)
		return
	}

	fmt.Println("Server inventory updated successfully.")
	fmt.Println("Server ID:", serverID)
	fmt.Println("Hostname:", serverInfo.Hostname)
	fmt.Println("IP Address:", serverInfo.IPAddress)
	fmt.Println("Operating System:", serverInfo.OperatingSystem)
	fmt.Println("RAM:", serverInfo.RAMGB, "GB")
	fmt.Println("CPU:", serverInfo.CPU)
	fmt.Println("Storage:", serverInfo.StorageGB, "GB")
	fmt.Println("GPU:", serverInfo.GPU)
	fmt.Println("Motherboard:", serverInfo.Motherboard)
}

func registerNetworkInterfaces() {
	fmt.Println("\nRegister Network Interfaces")
	fmt.Println("---")

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Configuration error:", err)
		return
	}

	serverInfo, err := inventory.Collect()
	if err != nil {
		fmt.Println("Inventory collection failed:", err)
		return
	}

	interfaces, err := network.Collect()
	if err != nil {
		fmt.Println("Network collection failed:", err)
		return
	}

	db, err := database.Connect(cfg)
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer db.Close()

	serverID, err := database.RegisterServer(db, serverInfo)
	if err != nil {
		fmt.Println("Server registration failed:", err)
		return
	}

	if err := database.UpsertNetworkInterfaces(db, serverID, interfaces); err != nil {
		fmt.Println("Network update failed:", err)
		return
	}

	fmt.Println("Network interfaces updated successfully.")

	for _, iface := range interfaces {
		fmt.Println()
		fmt.Println("Interface:", iface.Name)
		fmt.Println("MAC Address:", iface.MACAddress)
		fmt.Println("IP Address:", iface.IPAddress)
		fmt.Println("Network Type:", iface.NetworkType)
		fmt.Println("Speed:", iface.SpeedMbps, "Mbps")
	}
}
