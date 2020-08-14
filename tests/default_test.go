package exporter_tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"reflect"
	"gitlab.com/surfprace/cathal-go/exporter/parser"
)

func TestDefault(t *testing.T) {

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

	res1 := reflect.DeepEqual(mp["allegro"], allegro)
	res2 := reflect.DeepEqual(mp["lofar"], lofar)
	res3 := reflect.DeepEqual(mp["projectmine"], projectmine)
	res4 := reflect.DeepEqual(mp["sksp"], sksp)
	res5 := reflect.DeepEqual(mp["spexone"], spexone)

	fmt.Println("Map 1 Compare Result: ", res1)
	fmt.Println("Map 2 Compare Result: ", res2)
	fmt.Println("Map 3 Compare Result: ", res3)
	fmt.Println("Map 4 Compare Result: ", res4)
	fmt.Println("Map 5 Compare Result: ", res5)
}
