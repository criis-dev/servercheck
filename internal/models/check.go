package models

import "time"

type ServerCheckResult struct {
	ServerName string
	SSH        SSHResult
	Disk       []DiskResult
}

type SSHResult struct {
	Connected bool
	Latency   time.Duration
}
