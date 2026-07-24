package main

import (
    "fmt"
    "github.com/Kade710/ubuntu-sql-server/applications/go_agent/internal/system"
)

func main() {
    fmt.Println("---------------------------")
    fmt.Println(" Ubuntu SQL Server Go Agent")
    fmt.Println("---------------------------")
    fmt.Println()

    hostname, err := system.Hostname()
    if err != nil {
        fmt.Println("Hostname:", err)
    } else {
        fmt.Println("Hostname:", hostname)
    }

    fmt.Println("OperatingSystem:", system.OperatingSystem())
    fmt.Println("Architecture:", system.Architecture())
    fmt.Println()
    fmt.Println("Status : Placeholder")
    fmt.Println("Version: 0.1.0")
}