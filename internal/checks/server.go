package checks

import (
	"time"

	"github.com/cristianperen/servercheck/internal/models"
	"github.com/cristianperen/servercheck/internal/ssh"
)

func CheckServer(server models.Server) (models.ServerCheckResult, error) {
	start := time.Now()

	client, err := ssh.Connect(server)
	if err != nil {
		return models.ServerCheckResult{
			ServerName: server.Name,
			Status:     models.StatusDown,
			SSH: models.SSHResult{
				Connected: false,
				Latency:   time.Since(start),
			},
		}, nil
	}

	defer client.Close()

	latency := time.Since(start)

	output, err := ssh.RunCommand(client, "df -hl")
	if err != nil {
		return models.ServerCheckResult{
			ServerName: server.Name,
			Status:     models.StatusDown,
			SSH: models.SSHResult{
				Connected: true,
				Latency:   latency,
			},
		}, nil
	}

	diskResults, err := ParseDiskOutput(output)
	if err != nil {
		return models.ServerCheckResult{}, err
	}

	status := EvaluateServerStatus(diskResults)

	return models.ServerCheckResult{
		ServerName: server.Name,
		Status:     status,
		SSH: models.SSHResult{
			Connected: true,
			Latency:   latency,
		},
		Disk: diskResults,
	}, nil
}
