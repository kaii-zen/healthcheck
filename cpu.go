package main

import (
    "os"
    "fmt"
    "strings"
    "time"
    "github.com/shirou/gopsutil/cpu"
)

func check_cpu() {
    v, err := cpu.Percent(time.Second, true)
    if err != nil {
        panic(err.Error())
    }

    warn, crit := false, false
    var output []string

    // Sum of all percentages to calculate the average
    sum := 0.0
    for i, c := range v {
        // If a single core is at more than 90%
        // we want to trigger a warning.
        if c > 90 {
            warn = true
            output = append(output, fmt.Sprintf("[%v]: %.1f%%", i, c))
        }

        sum += c
    }

    avg := sum/float64(len(v))

    switch {
    case avg > 90:
        crit = true
    case avg > 80:
        warn = true
    }

    output = append([]string{statusString(warn, crit), fmt.Sprintf("Average: %.1f%%", avg)}, output...)
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
