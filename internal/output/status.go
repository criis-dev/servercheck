package output

import "github.com/cristianperen/servercheck/internal/models"

func formatStatus(status models.Status) string {
	switch status {
	case models.StatusOK:
		return "✓ OK"

	case models.StatusWarning:
		return "⚠ WARNING"

	case models.StatusCritical:
		return "✗ CRITICAL"

	case models.StatusDown:
		return "✗ DOWN"

	default:
		return string(status)
	}
}
