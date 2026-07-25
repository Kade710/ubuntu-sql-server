package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/config"
	"github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/database"
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

	fmt.Println("OperatingSystem:", system.OperatingSystem())
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
	fmt.Println("----------------------")

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
	fmt.Println()
	fmt.Println("Database Connection Test")
	fmt.Println("---")

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Configuration error:", err)
		return
	}

	db, err := database.Connect(cfg)
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer db.Close()

	fmt.Println("PostgreSQL connection successful")
}
