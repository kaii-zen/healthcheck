package main

import (
    "os"
    "fmt"
    "strings"
    "github.com/shirou/gopsutil/disk"
)

func check_disk() {
    partitions, err := disk.Partitions(true)
    if err != nil {
        fmt.Println("Error")
        os.Exit(2)
    }

    warn, crit := false, false
    var output []string

    for _, p := range partitions {
        u, err := disk.Usage(p.Mountpoint)
        if err != nil {
            panic(err.Error())
        }

        // Only include things that live in /dev and exclude things that always show at 100% (mounted images or CDs if you live in the 90s)
        if strings.HasPrefix(p.Device, "/dev") && p.Fstype != "cd9660" {
            switch {
            case u.UsedPercent > 90:
                crit = true
                fallthrough
            case u.UsedPercent > 80:
                warn = true

                output = append(output,
                                fmt.Sprintf("%v: %.1f%% (%.2f GiB Available/%.2f GiB Total)",
                                            p.Mountpoint, u.UsedPercent, toGiB(u.Free), toGiB(u.Total)))
            // Now same for inodes
            case u.InodesUsedPercent > 90:
                crit = true
                fallthrough
            case u.InodesUsedPercent > 80:
                warn = true
                output = append(output,
                                fmt.Sprintf("%v (inodes): %.1f%% (%.2f GiB Available/%.2f GiB Total)",
                                            p.Mountpoint, u.InodesUsedPercent, u.InodesFree, u.InodesTotal))
            }
        }
    }

    output = append([]string{statusString(warn, crit)}, output...)
    fmt.Println(strings.Join(output, " | "))

    switch {
    case crit:
        os.Exit(2)
    case warn:
        os.Exit(1)
    default:
        os.Exit(0)
    }
}
