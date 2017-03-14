package main

import (
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func check_cpu() {
	v, err := cpu.Percent(time.Second, true)
	if err != nil {
		panic(err.Error())
	}

	results := new(Results)

	// Sum of all percentages to calculate the average
	sum := 0.0
	for i, c := range v {
		results.Add(new(Result).Warn(90).Value(c).Output(fmt.Sprintf("[%v]: %.1f%%", i, c)))
		sum += c
	}

	avg := sum / float64(len(v))
	results.Add(new(Result).Warn(80).Crit(90).Value(avg).Output(fmt.Sprintf("Average: %.1f%%", avg)))

	fmt.Println(results)
	os.Exit(int(results.Status()))
}
