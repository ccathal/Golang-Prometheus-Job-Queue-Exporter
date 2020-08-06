package exporter_tests

import (
	"io/ioutil"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
	"gitlab.com/surfprace/cathal-go/src/exporter/parser"
)

func TestValidData(t *testing.T) {
	file, err := os.Open("test_data/slurm-out-valid.txt")
        if err != nil {
		t.Fatalf("Can not open test data: %v", err)
        }
        data, err := ioutil.ReadAll(file)

	mp := parser.ParseQueueMetrics(data)

	counter := 0
	for k := range mp {
		squeue_jobs := mp[k]
		for i := range squeue_jobs {
			counter += int(squeue_jobs[i])
		}
	}
	assert.Equal(t, counter, int(25))
}
