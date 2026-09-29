package output

import (
	"fmt"
	"strings"
)

func PrintTable(rows []ServerRow) {
	fmt.Printf(
		"%-20s %-20s %-18s %-18s %-15s\n",
		"SERVER",
		"SSH",
		"DISK",
		"INODES",
		"STATUS",
	)

	fmt.Println(strings.Repeat("-", 83))

	for _, row := range rows {
		fmt.Printf(
			"%-20s %-20s %-18s %-18s %-15s\n",
			row.Name,
			row.SSH,
			row.Disk,
			row.Inodes,
			row.Status,
		)
	}
}
