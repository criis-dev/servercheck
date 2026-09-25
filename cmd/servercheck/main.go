package main

import (
	"fmt"
	"os"

	"github.com/cristianperen/servercheck/internal/checks"
	"github.com/cristianperen/servercheck/internal/config"
	"github.com/cristianperen/servercheck/internal/ssh"
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
		// cfg, err := config.Load("configs/servers.yaml")
		// if err != nil {
		// 	fmt.Printf("Error loading configuration: %v\n", err)
		// 	return
		// }

		// for _, server := range cfg.Servers {
		// 	fmt.Printf("Connecting to %s...\n", server.Name)

		// 	client, err := ssh.Connect(server)
		// 	if err != nil {
		// 		fmt.Printf("  SSH error: %v\n", err)
		// 		continue
		// 	}
		// 	fmt.Printf("  SSH connetion OK")

		// 	client.Close()
		// }

		cfg, err := config.Load("configs/servers.yaml")
		if err != nil {
			fmt.Printf("Error loading configuration: %v\n", err)
			return
		}

		for _, server := range cfg.Servers {
			fmt.Printf("Connecting to %s...\n", server.Name)

			client, err := ssh.Connect(server)
			if err != nil {
				fmt.Printf("  SSH error: %v\n", err)
				continue
			}

			defer client.Close()

			output, err := ssh.RunCommand(client, "df -hl")
			results, err := checks.ParseDiskOutput(output)
			if err != nil {
				fmt.Printf("  Parse error: %v\n", err)
				continue
			}

			for _, result := range results {
				fmt.Printf(
					"  %s -> %d%% (%s)\n",
					result.MountPoint,
					result.UsagePercent,
					result.Filesystem,
				)
			}

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
