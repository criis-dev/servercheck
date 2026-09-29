package checks

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cristianperen/servercheck/internal/models"
)

func ParseInodeOutput(output string) ([]models.InodeResult, error) {
	lines := strings.Split(strings.TrimSpace(output), "\n")

	if len(lines) < 2 {
		return nil, fmt.Errorf("salida de df -il inválida")
	}

	var results []models.InodeResult

	for _, line := range lines[1:] {
		fields := strings.Fields(line)

		if len(fields) < 6 {
			continue
		}

		filesystem := fields[0]

		if !shouldCheckFilesystem(filesystem) {
			continue
		}

		usage := strings.TrimSuffix(fields[4], "%")

		usagePercent, err := strconv.Atoi(usage)
		if err != nil {
			return nil, fmt.Errorf("porcentaje de inodos inválido: %w", err)
		}

		results = append(results, models.InodeResult{
			Filesystem:   filesystem,
			UsagePercent: usagePercent,
			MountPoint:   fields[5],
			Status:       EvaluateDiskUsage(usagePercent),
		})
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no se encontraron filesystems de inodos")
	}

	return results, nil
}
