package models

type InodeResult struct {
	Filesystem   string
	UsagePercent int
	MountPoint   string
	Status       Status
}
