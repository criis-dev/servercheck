package checks

import "github.com/cristianperen/servercheck/internal/models"

func ParseDiskOutput(output string) ([]models.DiskResult, error) {
	results, err := parseFilesystemOutput(output)
	if err != nil {
		return nil, err
	}

	diskResults := make([]models.DiskResult, 0, len(results))

	for _, result := range results {
		diskResults = append(diskResults, models.DiskResult{
			Filesystem:   result.Filesystem,
			UsagePercent: result.UsagePercent,
			MountPoint:   result.MountPoint,
			Status:       result.Status,
		})
	}

	return diskResults, nil
}
