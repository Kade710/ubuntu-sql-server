package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/system"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("---------------------------")
        fmt.Println(" Ubuntu SQL Server Go Agent")
        fmt.Println("---------------------------")
        fmt.Println("1. Show System Information")
        fmt.Println("0. exit")
        fmt.Println("\nSelect an option")

        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice{

        case "1":
            showSystemInfo()

        case "0":
            fmt.Println("\nSee Ya!")
            return

        default:
            fmt.Println("\nTry Again!")
        }
    }
}

func showSystemInfo() {
    fmt.Println()

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
