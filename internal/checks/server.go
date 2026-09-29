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

	diskOutput, err := ssh.RunCommand(client, "df -hl")
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

	diskResults, err := ParseDiskOutput(diskOutput)
	if err != nil {
		return models.ServerCheckResult{}, err
	}

	inodeOutput, err := ssh.RunCommand(client, "df -il")
	if err != nil {
		return models.ServerCheckResult{}, err
	}

	inodeResults, err := ParseInodeOutput(inodeOutput)
	if err != nil {
		return models.ServerCheckResult{}, err
	}

	status := EvaluateServerStatus(diskResults, inodeResults)

	return models.ServerCheckResult{
		ServerName: server.Name,
		Status:     status,
		SSH: models.SSHResult{
			Connected: true,
			Latency:   latency,
		},
		Disk:   diskResults,
		Inodes: inodeResults,
	}, nil
}
