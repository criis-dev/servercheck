package main

import (
	"flag"
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
		runCheck(os.Args[2:])

	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}
}

func runCheck(args []string) {
	checkFlags := flag.NewFlagSet("check", flag.ExitOnError)

	server := checkFlags.String(
		"server",
		"",
		"check only the specified server",
	)

	verbose := checkFlags.Bool(
		"verbose",
		false,
		"show detailed check results",
	)

	checkFlags.Parse(args)

	cfg, err := config.Load("configs/servers.yaml")
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}

	rows := make([]output.ServerRow, 0, len(cfg.Servers))
	found := false

	for _, srv := range cfg.Servers {
		if *server != "" && srv.Name != *server {
			continue
		}

		found = true

		fmt.Printf("Connecting to %s...\n", srv.Name)

		result, err := checks.CheckServer(srv)
		if err != nil {
			fmt.Printf("  Error: %v\n", err)
			continue
		}

		rows = append(rows, output.BuildServerRow(result))

		if *verbose {
			output.PrintDetailedResult(result)
		}
	}

	if *server != "" && !found {
		fmt.Printf("Server not found: %s\n", *server)
		return
	}

	fmt.Println()
	output.PrintTable(rows)
}
