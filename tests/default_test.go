package exporter_tests

import (
	"io/ioutil"
	"os"
	"testing"
	"reflect"
	"gitlab.com/surfprace/cathal-go/exporter/parser"
)

func TestiDefault(t *testing.T) {

	allegro := map[string]int {
		"PENDING": 2,
		"RUNNING": 1,
		"COMPLETING" : 5,
		"COMPLETED": 1,
	}

        lofar := map[string]int {
                "PENDING": 1,
                "RUNNING": 1,
                "SUSPENDED": 1,
                "COMPLETING" : 1,
                "COMPLETED": 1,
        }

        sksp := map[string]int {
                "PENDING": 1,
                "RUNNING": 2,
                "COMPLETED": 2,
        }

        spexone := map[string]int {
                "RUNNING": 1,
        }

        projectmine := map[string]int {
                "PENDING": 1,
                "COMPLETING" : 3,
                "COMPLETED": 1,
        }

	file, err := os.Open("test_data/slurm-out-valid.txt")
        if err != nil {
		t.Fatalf("Can not open test data: %v", err)
        }
        data, err := ioutil.ReadAll(file)

	mp := parser.ParseQueueMetrics(data)

	reflect.DeepEqual(mp["allegro"], allegro)
	reflect.DeepEqual(mp["lofar"], lofar)
	reflect.DeepEqual(mp["projectmine"], projectmine)
	reflect.DeepEqual(mp["sksp"], sksp)
	reflect.DeepEqual(mp["spexone"], spexone)
}
