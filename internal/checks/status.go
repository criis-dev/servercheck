package checks

import "github.com/cristianperen/servercheck/internal/models"

func EvaluateDiskUsage(usagePercent int) models.Status {
	if usagePercent > 90 {
		return models.StatusCritical
	}

	if usagePercent >= 80 {
		return models.StatusWarning
	}

	return models.StatusOK
}

func EvaluateServerStatus(
	diskResults []models.DiskResult,
	inodeResults []models.InodeResult,
) models.Status {
	status := models.StatusOK

	for _, disk := range diskResults {
		if disk.Status == models.StatusCritical {
			return models.StatusCritical
		}

		if disk.Status == models.StatusWarning {
			status = models.StatusWarning
		}
	}

	for _, inode := range inodeResults {
		if inode.Status == models.StatusCritical {
			return models.StatusCritical
		}

		if inode.Status == models.StatusWarning {
			status = models.StatusWarning
		}
	}

	return status
}
