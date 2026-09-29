package main

import (
	"fmt"
	"os"

	"github.com/cristianperen/servercheck/internal/checks"
	"github.com/cristianperen/servercheck/internal/config"
	"github.com/cristianperen/servercheck/internal/output"
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
	fmt.Println("  servercheck check --server <name>")
	fmt.Println("  servercheck check --verbose")
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

		serverFilter := ""
		verbose := false

		for i := 2; i < len(os.Args); i++ {
			switch os.Args[i] {
			case "--verbose":
				verbose = true

			case "--server":
				if i+1 >= len(os.Args) {
					fmt.Println("Error: --server requires a server name")
					return
				}

				serverFilter = os.Args[i+1]
				i++
			}
		}

		rows := make([]output.ServerRow, 0, len(cfg.Servers))
		found := false

		for _, server := range cfg.Servers {
			if serverFilter != "" && server.Name != serverFilter {
				continue
			}

			found = true

			fmt.Printf("Connecting to %s...\n", server.Name)

			result, err := checks.CheckServer(server)
			if err != nil {
				fmt.Printf("  Error: %v\n", err)
				continue
			}

			rows = append(rows, output.BuildServerRow(result))

			if verbose {
				output.PrintDetailedResult(result)
			}
		}

		if serverFilter != "" && !found {
			fmt.Printf("Server not found: %s\n", serverFilter)
			return
		}

		fmt.Println()
		output.PrintTable(rows)

	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}
}
