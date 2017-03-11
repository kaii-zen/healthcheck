package main

import (
    "os"
    "fmt"
    "github.com/shirou/gopsutil/mem"
)

func check_mem() {
    v, err := mem.VirtualMemory()
    if err != nil {
        panic(err.Error())
    }

    warn, crit := false, false

    switch {
    case v.UsedPercent > 95:
        crit = true
    case v.UsedPercent > 80:
        warn = true
    }


    fmt.Printf( "%s | Total: %.2f GiB | Available: %.2f GiB | Used: %.1f%%\n",
                statusString(warn, crit),
                toGiB(v.Total),
                toGiB(v.Available),
                v.UsedPercent )

    switch {
    case crit:
        os.Exit(2)
    case warn:
        os.Exit(1)
    default:
        os.Exit(0)
    }
}
