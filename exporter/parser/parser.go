package parser

import (
	"strings"
)

func ParseQueueMetrics(input []byte) map[string]map[string]int {

	squeue_info :=  make(map[string]map[string]int)

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		splitted := strings.Split(line, ",")
		if len(splitted) == 12 {
			project := splitted[2]
			state := splitted[9]
			if len(state) > 0 && !strings.ContainsAny(state, " ") {
				if squeue_info[project] == nil {
					copy_map := make(map[string]int)
					squeue_info[project] = copy_map
				}
				if _, ok := squeue_info[project][state]; !ok {
					squeue_info[project][state] = 0
				}
				squeue_info[project][state] += 1
			}
		}
	}
	return squeue_info
}
