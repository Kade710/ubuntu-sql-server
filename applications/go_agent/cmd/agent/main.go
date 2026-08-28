package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/alerts"
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
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--refresh":
			if err := refreshServerData(); err != nil {
				fmt.Println("Server refresh failed:", err)
				os.Exit(1)
			}
			return

		case "--help", "-h":
			fmt.Println("Usage:")
			fmt.Println("	./agent				Start interactive mode")
			fmt.Println("	./agent --refresh	Refresh all server data")
			return

		default:
			fmt.Println("That's NOT the right option, please try again:", os.Args[1])
			fmt.Println("Use --help for available options.")
			os.Exit(2)
		}
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
		fmt.Println("0. Exit")
		fmt.Println("\nSelect an option: ")

		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nInput closed.")
			return
		}

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
			if err := refreshServerData(); err != nil {
				fmt.Println("\nServer refresh failed:", err)
			}

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

	serverID, serverInfo, err := registerServerInventoryCore()
	if err != nil {
		fmt.Println("Server inventory update failed:", err)
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

func registerServerInventoryCore() (int, inventory.Server, error) {
	cfg, err := config.Load()
	if err != nil {
		return 0, inventory.Server{}, fmt.Errorf("load configuration: %w", err)
	}

	serverInfo, err := inventory.Collect()
	if err != nil {
		return 0, inventory.Server{}, fmt.Errorf("collect inventory: %w", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		return 0, inventory.Server{}, fmt.Errorf("connect to database: %w", err)
	}
	defer db.Close()

	serverID, err := database.RegisterServer(db, serverInfo)
	if err != nil {
		return 0, inventory.Server{}, fmt.Errorf("register server: %w", err)
	}

	return serverID, serverInfo, nil
}

func registerNetworkInterfaces() {
	fmt.Println("\nRegister Network Interfaces")
	fmt.Println("---")

	interfaces, err := registerNetworkInterfacesCore()
	if err != nil {
		fmt.Println("Network interface update failed:", err)
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

func registerNetworkInterfacesCore() ([]network.Interface, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("load configuration: %w", err)
	}

	serverInfo, err := inventory.Collect()
	if err != nil {
		return nil, fmt.Errorf("collect inventory: %w", err)
	}

	interfaces, err := network.Collect()
	if err != nil {
		return nil, fmt.Errorf("collect network interfaces: %w", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}
	defer db.Close()

	serverID, err := database.RegisterServer(db, serverInfo)
	if err != nil {
		return nil, fmt.Errorf("register server: %w", err)
	}

	if err := database.UpsertNetworkInterfaces(db, serverID, interfaces); err != nil {
		return nil, fmt.Errorf("update network interfaces: %w", err)
	}

	return interfaces, nil
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
	if err != nil {
		fmt.Println("Input failed:", err)
		return
	}

	description := readInput(reader, "Description: ")
	if err != nil {
		fmt.Println("Input failed:", err)
		return
	}
	
	performedBy := readInput(reader, "Performed by: ")
	if err != nil {
		fmt.Println("Input failed:", err)
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
		return "", fmt.Errorf("read input: %w", err)
	}

	return strings.TrimSpace(value), nil
}

func readRequiredInput(reader *bufio.Reader, prompt string) string {
	for {
		value := readInput(reader, prompt)
		if value != nil {
			return "", err
		}

		if value != "" {
			return value, nil
		}

		fmt.Println("This field is required.")
	}
}

func registerOperatingSystem() {
	fmt.Println("\nRegister Operating System")
	fmt.Println("---")

	osID, serverID, osDetails, err := registerOperatingSystemCore()
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

func registerOperatingSystemCore() (int, int, osinfo.Info, error) {
	cfg, err := config.Load()
	if err != nil {
		return 0, 0, osinfo.Info{}, fmt.Errorf("load configuration: %w", err)
	}

	serverInfo, err := inventory.Collect()
	if err != nil {
		return 0, 0, osinfo.Info{}, fmt.Errorf("collect inventory: %w", err)
	}

	osDetails := osinfo.Collect()

	db, err := database.Connect(cfg)
	if err != nil {
		return 0, 0, osinfo.Info{}, fmt.Errorf("connect to database: %w", err)
	}
	defer db.Close()

	serverID, err := database.RegisterServer(db, serverInfo)
	if err != nil {
		return 0, 0, osinfo.Info{}, fmt.Errorf("register server: %w", err)
	}

	osID, err := database.UpsertOperatingSystem(
		db,
		serverID,
		osDetails,
	)
	if err != nil {
		return 0, 0, osinfo.Info{}, fmt.Errorf("update operating system: %w", err)
	}

	return osID, serverID, osDetails, nil
}

func registerHardwareComponents() {
	fmt.Println("\nRegister Hardware Components")
	fmt.Println("---")

	components, err := registerHardwareComponentsCore()
	if err != nil {
		fmt.Println("Hardware component update failed:", err)
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

func registerHardwareComponentsCore() ([]hardware.Component, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("load configuration: %w", err)
	}

	serverInfo, err := inventory.Collect()
	if err != nil {
		return nil, fmt.Errorf("collect inventory: %w", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}
	defer db.Close()

	serverID, err := database.RegisterServer(db, serverInfo)
	if err != nil {
		return nil, fmt.Errorf("register server: %w", err)
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
		return nil, fmt.Errorf("update hardware components: %w", err)
	}

	return components, nil
}

func showSystemHealth() {
	fmt.Println("\nSystem Health")
	fmt.Println("---")

	status, healthCheckID, err := showSystemHealthCore()
	if err != nil {
		fmt.Println("Health check failed:", err)
		return
	}

	fmt.Printf("1-Minute Load Average: %.2f\n", status.Load1)
	fmt.Printf("Memory Usage: %.2f%%\n", status.MemoryPercent)
	fmt.Printf("Disk Usage: %.2f%%\n", status.DiskPercent)
	fmt.Printf("Uptime: %.2f hours\n", status.UptimeHours)
	fmt.Println("Overall Status:", status.Overall)

	fmt.Println("Health check saved successfully.")
	fmt.Println("Health check ID:", healthCheckID)
}

func showSystemHealthCore() (health.Status, int, error) {
	status, err := health.Collect()
	if err != nil {
		return health.Status{}, 0, fmt.Errorf("collect system health: %w", err)
	}

	cfg, err := config.Load()
	if err != nil {
		return status, 0, fmt.Errorf("load configuration: %w", err)
	}

	serverInfo, err := inventory.Collect()
	if err != nil {
		return status, 0, fmt.Errorf("collect inventory: %w", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		return status, 0, fmt.Errorf("connect to database: %w", err)
	}
	defer db.Close()

	serverID, err := database.RegisterServer(db, serverInfo)
	if err != nil {
		return status, 0, fmt.Errorf("register server: %w", err)
	}

	previousStatus := ""

	records, err := database.GetRecentHealthChecks(db, serverID, 1)
	if err != nil {
		fmt.Println("Previous health status unavailable:", err)
	} else if len(records) > 0 {
		previousStatus = records[0].OverallStatus
	}

	healthCheckID, err := database.AddHealthCheck(db, serverID, status)
	if err != nil {
		return status, 0, fmt.Errorf("save health check: %w", err)
	}

	if previousStatus != status.Overall {
		title := "U-Server Health Alert"

		var message string

		switch status.Overall {
		case "WARNING":
			message = fmt.Sprintf(
				"U-Server WARNING\nLoad: %.2f\nMemory: %.2f%%\nDisk: %.2f%%",
				status.Load1,
				status.MemoryPercent,
				status.DiskPercent,
			)

		case "CRITICAL":
			message = fmt.Sprintf(
				"U-Server CRITICAL\nLoad: %.2f\nMemory: %.2f%%\nDisk: %.2f%%",
				status.Load1,
				status.MemoryPercent,
				status.DiskPercent,
			)

		case "HEALTHY":
			if previousStatus == "WARNING" || previousStatus == "CRITICAL" {
				title = "U-Server Recovered"

				message = fmt.Sprintf(
					"U-Server has recovered.\nLoad: %.2f\nMemory: %.2f%%\nDisk: %.2f%%",
					status.Load1,
					status.MemoryPercent,
					status.DiskPercent,
				)
			}
		}

		if message != "" {
			alertID, err := database.AddAlertEvent(
				db,
				serverID,
				previousStatus,
				status.Overall,
				title,
				message,
			)

			if err != nil {
				fmt.Println("Alert event not saved:", err)
			} else {
				fmt.Println("Alert event saved successfully.")
				fmt.Println("Alert event ID:", alertID)
			}

			if err := alerts.Send(cfg.NTFYTopic, title, message); err != nil {
				fmt.Println("Notification failed:", err)
			} else {
				fmt.Println("Notification sent successfully.")
			}
		}
	}

	return status, healthCheckID, nil
}

func viewRecentHealthChecks() {
	fmt.Println("\nRecent Health Checks")
	fmt.Println("---")

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
		fmt.Println("No health checks found.")
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

func refreshServerData() error {
	fmt.Println("\nRefresh Server Data")
	fmt.Println("---")

	fmt.Println("\n[1/5] Updating server inventory....")
	if _, _, err := registerServerInventoryCore(); err != nil {
		return fmt.Errorf("server inventory refresh failed: %w", err)
	}
	fmt.Println("Server inventory updated successfully.")

	fmt.Println("\n[2/5] Updating operating system....")
	if _, _, _, err := registerOperatingSystemCore(); err != nil {
		return fmt.Errorf("operating system refresh failed: %w", err)
	}
	fmt.Println("Operating system updated successfully.")

	fmt.Println("\n[3/5] Updating hardware components....")
	if _, err := registerHardwareComponentsCore(); err != nil {
		return fmt.Errorf("hardware refresh failed: %w", err)
	}
	fmt.Println("Hardware components updated successfully.")

	fmt.Println("\n[4/5] Updating network interfaces....")
	if _, err := registerNetworkInterfacesCore(); err != nil {
		return fmt.Errorf("network refresh failed: %w", err)
	}
	fmt.Println("Network interfaces updated successfully.")

	fmt.Println("\n[5/5] Recording system health....")
	status, healthCheckID, err := showSystemHealthCore()
	if err != nil {
		return fmt.Errorf("health refresh failed: %w", err)
	}

	fmt.Printf("1-Minute Load Average: %.2f\n", status.Load1)
	fmt.Printf("Memory Usage: %.2f%%\n", status.MemoryPercent)
	fmt.Printf("Disk Usage: %.2f%%\n", status.DiskPercent)
	fmt.Printf("Uptime: %.2f hours\n", status.UptimeHours)
	fmt.Println("Overall Status:", status.Overall)
	fmt.Println("Health check saved successfully.")
	fmt.Println("Health check ID:", healthCheckID)

	fmt.Println("\nServer refresh completed successfully.")

	return nil
}
