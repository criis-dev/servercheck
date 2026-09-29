package checks

type Status string

const (
	StatusOK       Status = "OK"
	StatusWarning  Status = "WARNING"
	StatusCritical Status = "CRITICAL"
)

func EvaluateDiskUsage(usagePercent int) Status {
	if usagePercent > 90 {
		return StatusCritical
	}

	if usagePercent >= 80 {
		return StatusWarning
	}

	return StatusOK
}
