package checks

import "github.com/cristianperen/servercheck/internal/models"

func ParseInodeOutput(output string) ([]models.InodeResult, error) {
	results, err := parseFilesystemOutput(output)
	if err != nil {
		return nil, err
	}

	inodeResults := make([]models.InodeResult, 0, len(results))

	for _, result := range results {
		inodeResults = append(inodeResults, models.InodeResult{
			Filesystem:   result.Filesystem,
			UsagePercent: result.UsagePercent,
			MountPoint:   result.MountPoint,
			Status:       result.Status,
		})
	}

	return inodeResults, nil
}
