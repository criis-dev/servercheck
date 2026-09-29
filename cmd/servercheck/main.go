package main

import (
	"fmt"
	"os"

	"github.com/cristianperen/servercheck/internal/checks"
	"github.com/cristianperen/servercheck/internal/config"
	"github.com/cristianperen/servercheck/internal/models"
)

const version = "0.1.0"

func showVersion() {
	fmt.Printf("ServerCheck v%s\n", version)
}

func showHelp() {
	fmt.Println("ServerCheck - SSH server health checker")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  servercheck version")
	fmt.Println("  servercheck check")
}

func checkServers() {
	cfg, err := config.Load("configs/servers.yaml")
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}

	fmt.Println("Checking servers...")
	fmt.Println()

	for _, server := range cfg.Servers {
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
		cfg, err := config.Load("configs/servers.yaml")
		if err != nil {
			fmt.Printf("Error loading configuration: %v\n", err)
			return
		}

		for _, server := range cfg.Servers {
			fmt.Printf("Connecting to %s...\n", server.Name)

			// start := time.Now()

			// latency := time.Since(start)

			result, err := checks.CheckServer(server)
			if err != nil {
				fmt.Printf("  Error: %v\n", err)
				continue
			}

			if result.Status == models.StatusDown {
				fmt.Printf("  SSH -> DOWN\n")
				continue
			}

			for _, disk := range result.Disk {

				fmt.Printf(
					"  %s -> %d%% (%s) [%s]\n",
					disk.MountPoint,
					disk.UsagePercent,
					disk.Filesystem,
					disk.Status,
				)
			}

			fmt.Printf("  Server status %v\n", result.Status)

			if err != nil {
				fmt.Printf("  Command error: %v\n", err)
				continue
			}

			// fmt.Println(output)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}
}
