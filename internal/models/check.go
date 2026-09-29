package models

import "time"

type ServerCheckResult struct {
	ServerName string
	Status     Status
	SSH        SSHResult
	Disk       []DiskResult
	Inodes     []InodeResult
}

type SSHResult struct {
	Connected bool
	Latency   time.Duration
}
