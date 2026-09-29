package output

import (
	"fmt"

	"github.com/cristianperen/servercheck/internal/models"
)

type ServerRow struct {
	Name   string
	SSH    string
	Disk   string
	Inodes string
	Status string
}

func BuildServerRow(result models.ServerCheckResult) ServerRow {
	return ServerRow{
		Name:   result.ServerName,
		SSH:    formatSSH(result.SSH),
		Disk:   formatUsage(result.Disk),
		Inodes: formatInodes(result.Inodes),
		Status: formatStatus(result.Status),
	}
}

func formatSSH(result models.SSHResult) string {
	if !result.Connected {
		return "✗ DOWN"
	}

	return fmt.Sprintf("✓ %s", result.Latency)
}

func formatUsage(results []models.DiskResult) string {
	if len(results) == 0 {
		return "—"
	}

	status := models.StatusOK
	maxUsage := 0

	for _, result := range results {
		if result.UsagePercent > maxUsage {
			maxUsage = result.UsagePercent
		}

		if result.Status == models.StatusCritical {
			status = models.StatusCritical
		} else if result.Status == models.StatusWarning && status != models.StatusCritical {
			status = models.StatusWarning
		}
	}

	return fmt.Sprintf("%s %d%%", formatStatus(status), maxUsage)
}

func formatInodes(results []models.InodeResult) string {
	if len(results) == 0 {
		return "—"
	}

	status := models.StatusOK
	maxUsage := 0

	for _, result := range results {
		if result.UsagePercent > maxUsage {
			maxUsage = result.UsagePercent
		}

		if result.Status == models.StatusCritical {
			status = models.StatusCritical
		} else if result.Status == models.StatusWarning && status != models.StatusCritical {
			status = models.StatusWarning
		}
	}

	return fmt.Sprintf("%s %d%%", formatStatus(status), maxUsage)
}
