package models

import "time"

type ServerCheckResult struct {
	ServerName string
	Status     Status
	SSH        SSHResult
	Disk       []DiskResult
}

type SSHResult struct {
	Connected bool
	Latency   time.Duration
}
