package checks

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cristianperen/servercheck/internal/models"
)

func ParseDiskOutput(output string) ([]models.DiskResult, error) {
	lines := strings.Split(strings.TrimSpace(output), "\n")

	if len(lines) < 2 {
		return nil, fmt.Errorf("salida de df inválida")
	}

	var results []models.DiskResult

	for _, line := range lines[1:] {
		fields := strings.Fields(line)

		if len(fields) < 6 {
			continue
		}

		if !shouldCheckFilesystem(fields[0]) {
			continue
		}

		filesystem := fields[0]
		usage := strings.TrimSuffix(fields[4], "%")

		usagePercent, err := strconv.Atoi(usage)
		if err != nil {
			return nil, fmt.Errorf("porcentaje inválido: %w", err)
		}

		results = append(results, models.DiskResult{
			Filesystem:   filesystem,
			UsagePercent: usagePercent,
			MountPoint:   fields[5],
			Status:       EvaluateDiskUsage(usagePercent),
		})
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no se encontraron filesystems")
	}

	return results, nil
}

func shouldCheckFilesystem(filesystem string) bool {
	if strings.HasPrefix(filesystem, "tmpfs") {
		return false
	}

	if strings.HasPrefix(filesystem, "devtmpfs") {
		return false
	}

	return true
}
