package main

import (
	"fmt"
	"os"

	"github.com/cristianperen/servercheck/internal/models"
)

const vesion = "0.1.0"

func showVersion() {
    fmt.Println("ServerCheck v0.1.0")
}

func showHelp() {
    fmt.Println("ServerCheck - SSH server health checker")
    fmt.Println("")
    fmt.Println("Usage:")
    fmt.Println(" servercheck version")
    fmt.Println(" servercheck check")
}

func checkServers() {
    servers := []models.Server{
        {
            Name: "web-01",
            Host: "192.168.1.10",
            Port: 22,
            User: "root",
        },
        {
            Name: "web-02",
            Host: "192.168.1.11",
            Port: 22,
            User: "ubuntu",
        },
    }

    fmt.Println("Checking servers...")
    fmt.Println()

    for _, server := range servers {
        fmt.Printf("• %s (%s:%d)\n", server.Name, server.Host, server.Port)
    }
}

func main() {
    if len(os.Args) < 2 {
        showHelp()
        return
    }

    command := os.Args[1]

    switch command {
        case "version":
            showVersion()

        case "check":
            fmt.Println("Checking servers...")

        default:
            fmt.Printf("Unknown command: %s\n", command)
            showHelp()
    }
}