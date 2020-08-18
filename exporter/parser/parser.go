package parser

import (
	"strings"
)

func ParseQueueMetrics(input []byte) map[string]map[string]int {

	// nested map returned to main exporter
	squeue_info :=  make(map[string]map[string]int)

	// loop through array of squeue data
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		splitted := strings.Split(line, ",")
		// output length of squeue is 12
		if len(splitted) == 12 {
			project := splitted[2]
			state := splitted[9]
			// ensure job type is only 1 word
			if len(state) > 0 && !strings.ContainsAny(state, " ") {
				// if project key does not exist, initialise key with empty map value
				if squeue_info[project] == nil {
					copy_map := make(map[string]int)
					squeue_info[project] = copy_map
				}
				// if job type does not exist yet in value map, initialise with value 0
				if _, ok := squeue_info[project][state]; !ok {
					squeue_info[project][state] = 0
				}
				// increment job type value by 1
				squeue_info[project][state] += 1
			}
		}
	}
	// return nested map
	return squeue_info
}
