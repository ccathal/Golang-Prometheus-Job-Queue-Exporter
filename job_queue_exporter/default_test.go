package exporter_tests

import (
	"io/ioutil"
	"os"
	"testing"
	"reflect"
	"gitlab.com/surfprace/cathal-go/job_queue_exporter"
)

func TestiDefault(t *testing.T) {

	allegro := map[string]int {
		"PENDING": 2,
		"RUNNING": 1,
		"SUSPENDED": 0,
		"CANCELLED": 0,
		"COMPLETING" : 5,
		"COMPLETED": 1,
		"CONFIGURING" : 0,
		"FAILED": 0,
		"TIMEOUT": 0,
		"PREEMPTED": 0,
		"NODE_FAIL": 0,
	}

        lofar := map[string]int {
                "PENDING": 1,
                "RUNNING": 1,
                "SUSPENDED": 1,
                "CANCELLED": 0,
                "COMPLETING" : 1,
                "COMPLETED": 1,
                "CONFIGURING" : 0,
                "FAILED": 0,
                "TIMEOUT": 0,
                "PREEMPTED": 0,
                "NODE_FAIL": 0,
        }

        sksp := map[string]int {
                "PENDING": 1,
                "RUNNING": 2,
                "SUSPENDED": 0,
                "CANCELLED": 0,
                "COMPLETING" : 0,
                "COMPLETED": 2,
                "CONFIGURING" : 0,
                "FAILED": 0,
                "TIMEOUT": 0,
                "PREEMPTED": 0,
                "NODE_FAIL": 0,
        }

        spexone := map[string]int {
                "PENDING": 0,
                "RUNNING": 1,
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

        projectmine := map[string]int {
                "PENDING": 1,
                "RUNNING": 0,
                "SUSPENDED": 0,
                "CANCELLED": 0,
                "COMPLETING" : 3,
                "COMPLETED": 1,
                "CONFIGURING" : 0,
                "FAILED": 0,
                "TIMEOUT": 0,
                "PREEMPTED": 0,
                "NODE_FAIL": 0,
        }



	file, err := os.Open("test_data/slurm-out-valid.txt")
        if err != nil {
		t.Fatalf("Can not open test data: %v", err)
        }
        data, err := ioutil.ReadAll(file)

	mp := ParseQueueMetrics(data)

	reflect.DeepEqual(mp["allegro"], allegro)
	reflect.DeepEqual(mp["lofar"], lofar)
	reflect.DeepEqual(mp["projectmine"], projectmine)
	reflect.DeepEqual(mp["sksp"], sksp)
	reflect.DeepEqual(mp["spexone"], spexone)
}
