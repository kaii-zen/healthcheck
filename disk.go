package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/shirou/gopsutil/disk"
)

func check_disk() {
	partitions, err := disk.Partitions(true)
	if err != nil {
		panic(err.Error())
	}

	resultz := new(Results)

	for _, p := range partitions {
		u, err := disk.Usage(p.Mountpoint)
		if err != nil {
			panic(err.Error())
		}

		// Only include things that live in /dev and exclude things that always show at 100% (mounted images or CDs if you live in the 90s)
		if strings.HasPrefix(p.Device, "/dev") && p.Fstype != "cd9660" {
			resultz.Add(new(Result).Value(u.UsedPercent).Crit(90).Warn(80).
				Output(fmt.Sprintf("%v: %.1f%% (%.2f GiB Available/%.2f GiB Total)",
					p.Mountpoint, u.UsedPercent, toGiB(u.Free), toGiB(u.Total))))

			resultz.Add(new(Result).Value(u.InodesUsedPercent).Crit(90).Warn(80).
				Output(fmt.Sprintf("%v (inodes): %.1f%% (%.2f Available/%.2f Total)",
					p.Mountpoint, u.InodesUsedPercent, u.InodesFree, u.InodesTotal)))
		}
	}

	fmt.Println(resultz)
	os.Exit(int(resultz.Status()))

}
