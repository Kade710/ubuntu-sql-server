package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/config"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/database"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/hardware"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/health"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/inventory"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/maintenance"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/network"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/osinfo"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/system"
)

func main() {
		if len(os.Args) > 1 && os.Args[1] == "--refresh" {
			refreshServerData()
			return
		}
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
		fmt.Println("6. Add Maintenance Log")
		fmt.Println("7. Register Operating System")
		fmt.Println("8. Register Hardware Components")
		fmt.Println("9. Show System Health")
		fmt.Println("10. View Recent Health Checks")
		fmt.Println("11. Refresh Server Data")
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

		case "6":
			addMaintenanceLog(reader)

		case "7":
			registerOperatingSystem()

		case "8":
			registerHardwareComponents()

		case "9":
			showSystemHealth()

		case "10":
			viewRecentHealthChecks()

		case "11":
			refreshServerData()

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

func addMaintenanceLog(reader *bufio.Reader) {
	fmt.Println("\nAdd Maintenance Log")
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

	action := readRequiredInput(reader, "Action: ")
	description := readInput(reader, "Description: ")
	performedBy := readInput(reader, "Performed by: ")

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

	logEntry := maintenance.Log{
		Action:      action,
		Description: description,
		PerformedBy: performedBy,
	}

	logID, err := database.AddMaintenanceLog(db, serverID, logEntry)
	if err != nil {
		fmt.Println("Maintenance log failed:", err)
		return
	}

	fmt.Println()
	fmt.Println("Maintenance log saved successfully.")
	fmt.Println("Log ID:", logID)
	fmt.Println("Server ID:", serverID)
	fmt.Println("Action:", logEntry.Action)
	fmt.Println("Description:", logEntry.Description)
	fmt.Println("Performed by:", logEntry.PerformedBy)
}

func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)

	value, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(value)
}

func readRequiredInput(reader *bufio.Reader, prompt string) string {
	for {
		value := readInput(reader, prompt)
		if value != "" {
			return value
		}

		fmt.Println("This field is required.")
	}
}

func registerOperatingSystem() {
	fmt.Println("\nRegister Operating System")
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

	osDetails := osinfo.Collect()

	db, err := database.Connect(cfg)
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer db.Close()

	serverID, err := database.RegisterServer(db, serverInfo)
	if err != nil {
		fmt.Println("Server registration failed", err)
		return
	}

	osID, err := database.UpsertOperatingSystem(
		db,
		serverID,
		osDetails,
	)
	if err != nil {
		fmt.Println("Operating system update failed:", err)
		return
	}

	fmt.Println("Operating system updated successfully.")
	fmt.Println("OS ID:", osID)
	fmt.Println("Server ID:", serverID)
	fmt.Println("Distribution:", osDetails.Distribution)
	fmt.Println("Version:", osDetails.Version)
	fmt.Println("Kernel:", osDetails.Kernel)
	fmt.Println("Architecture:", osDetails.Architecture)
}

func registerHardwareComponents() {
	fmt.Println("\nRegister Hardware Components")
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
		fmt.Println("Server registration failed:", err)
		return
	}

	hardwareInfo := hardware.Collect()

	components := hardware.Components(
		hardwareInfo,
		serverInfo.RAMGB,
	)

	if err := database.UpsertHardwareComponents(
		db,
		serverID,
		components,
	); err != nil {
		fmt.Println("Hardware update failed:", err)
		return
	}

	fmt.Println("Hardware components updated successfully.")

	for _, component := range components {
		fmt.Println()
		fmt.Println("Component:", component.Type)

		if component.Manufacturer != "" {
			fmt.Println("Manufacturer:", component.Manufacturer)
		}

		if component.Model != "" {
			fmt.Println("Model:", component.Model)
		}

		if component.Specification != "" {
			fmt.Println("Specification:", component.Specification)
		}
	}
}

func showSystemHealth() {
	fmt.Println("\nSystem Health")
	fmt.Println("---")

	status, err := health.Collect()
	if err != nil {
		fmt.Println("Health collection faled:", err)
		return
	}

	fmt.Printf("1-Minute Load Average: %.2f\n", status.Load1)
	fmt.Printf("Memory Usage: %.2f%%\n", status.MemoryPercent)
	fmt.Printf("Disk Usage: %.2f%%\n", status.DiskPercent)
	fmt.Printf("Uptime: %.2f hours\n", status.UptimeHours)
	fmt.Println("Overall Status:", status.Overall)

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Health reading not saved:", err)
		return
	}

	serverInfo, err := inventory.Collect()
	if err != nil {
		fmt.Println("Health reading not saved:", err)
		return
	}

	db, err := database.Connect(cfg)
	if err != nil {
		fmt.Println("Health reading not saved:", err)
		return
	}
	defer db.Close()

	serverID, err := database.RegisterServer(db, serverInfo)
	if err != nil {
		fmt.Println("Health reading not saved:", err)
		return
	}

	healthCheckID, err := database.AddHealthCheck(db, serverID, status)
	if err != nil {
		fmt.Println("Health reading not saved:", err)
		return
	}

	fmt.Println("Health check saved successfully.")
	fmt.Println("Health check ID:", healthCheckID)
}

func viewRecentHealthChecks() {
	fmt.Printf("\nRecent Health Checks")
	fmt.Print("---")

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Configuration Error:", err)
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
		fmt.Println("Server registration failed:", err)
		return
	}

	records, err := database.GetRecentHealthChecks(db, serverID, 5)
	if err != nil {
		fmt.Println("Health history failed:", err)
		return
	}

	if len(records) == 0 {
		fmt.Println("No helth check found.")
		return
	}

	for _, record := range records {
		fmt.Println()
		fmt.Println("Health Check ID:", record.ID)
		fmt.Printf("Load Average: %.2f\n", record.LoadAverage)
		fmt.Printf("Memory Percent: %.2f%%\n", record.MemoryPercent)
		fmt.Printf("Disk Percent: %.2f%%\n", record.DiskPercent)
		fmt.Printf("Uptime Hours: %.2f hours\n", record.UptimeHours)
		fmt.Println("Overall Status:", record.OverallStatus)
		fmt.Println("Created At:", record.CreatedAt)
	}
}

func refreshServerData() {
	fmt.Println("\nRefresh Server Data")
	fmt.Println("---")

	fmt.Println("\n[1/5] Updating server inventory....")
	registerServerInventory()

	fmt.Println("\n[2/5] Updating operating system....")
	registerOperatingSystem()

	fmt.Println("\n[3/5] Updating hardware components....")
	registerHardwareComponents()

	fmt.Println("\n[4/5] Updating network interfaces....")
	registerNetworkInterfaces()

	fmt.Println("\n[5/5] Recording system health....")
	showSystemHealth()

	fmt.Println("\nServer refresh completed!!!")
}
