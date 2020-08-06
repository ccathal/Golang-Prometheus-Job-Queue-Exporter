package main

import (
	"strings"
)

func ParseQueueMetrics(input []byte) map[string]map[string]int {

	squeue_info :=  make(map[string]map[string]int)
	states := map[string]int{
		"PENDING": 0,
		"RUNNING": 0,
		"SUSPENDED": 0,
		"CANCELLED": 0,
		"COMPLETING" : 0,
		"COMPLETED": 0,
		"CONFIGURING" : 0,
		"FAILED": 0,
		"TIMEOUT": 0,
		"PREEMPTED": 0,
		"NODE_FAIL": 0,
	}

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		splitted := strings.Split(line, ",")
		if len(splitted) == 12 {
			project := splitted[2]
			state := splitted[9]
			if _, ok := states[state]; ok {
				if squeue_info[project] == nil {
					copy_map := make(map[string]int)
					for key,value := range states {
						copy_map[key] = value
					}
					squeue_info[project] = copy_map
				}
				squeue_info[project][state] += 1
			}
		}
	}
	return squeue_info
}
