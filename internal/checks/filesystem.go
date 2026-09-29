package checks

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cristianperen/servercheck/internal/models"
)

type filesystemResult struct {
	Filesystem   string
	UsagePercent int
	MountPoint   string
	Status       models.Status
}

func parseFilesystemOutput(output string) ([]filesystemResult, error) {
	lines := strings.Split(strings.TrimSpace(output), "\n")

	if len(lines) < 2 {
		return nil, fmt.Errorf("salida de filesystem inválida")
	}

	var results []filesystemResult

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
			return nil, fmt.Errorf("porcentaje inválido: %w", err)
		}

		results = append(results, filesystemResult{
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
	return !strings.HasPrefix(filesystem, "tmpfs")
}
