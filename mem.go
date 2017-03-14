package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/mem"
)

func check_mem() {
	v, err := mem.VirtualMemory()
	if err != nil {
		panic(err.Error())
	}

	results := new(Results).
		Add(
			new(Result).
				Value(v.UsedPercent).
				Crit(95).
				Warn(90).
				Output(
					fmt.Sprintf(
						"Total: %.2f GiB | Available: %.2f GiB | Used: %.1f%%",
						toGiB(v.Total),
						toGiB(v.Available),
						v.UsedPercent)))

	fmt.Println(results)
	os.Exit(int(results.Status()))
}
