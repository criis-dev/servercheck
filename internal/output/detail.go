package output

import (
	"fmt"

	"github.com/cristianperen/servercheck/internal/models"
)

func PrintDetailedResult(result models.ServerCheckResult) {
	fmt.Println("  Disk:")

	for _, disk := range result.Disk {
		fmt.Printf(
			"    %s -> %d%% (%s) [%s]\n",
			disk.MountPoint,
			disk.UsagePercent,
			disk.Filesystem,
			formatStatus(disk.Status),
		)
	}

	fmt.Println("  Inodes:")

	for _, inode := range result.Inodes {
		fmt.Printf(
			"    %s -> %d%% (%s) [%s]\n",
			inode.MountPoint,
			inode.UsagePercent,
			inode.Filesystem,
			formatStatus(inode.Status),
		)

	}

	fmt.Println()
}
