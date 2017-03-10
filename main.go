package main

import (
    "os"
    "fmt"
)

func usage() {
    fmt.Println("Usage: healthcheck [cpu|mem|disk]")
}

func main() {
    if len(os.Args) == 1 {
        usage()
        os.Exit(0)
    }

    switch os.Args[1] {
    case "cpu":
        check_cpu()
    case "mem", "memory":
        check_mem()
    case "disk":
        check_disk()
    default:
        usage()
    }
}

func statusString(warn, crit bool) string {
    switch {
    case crit:
        return "CRIT"
    case warn:
        return "WARN"
    default:
        return "OK"
    }
}
