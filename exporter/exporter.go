package main

import (
        "strings"
        "flag"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promhttp"
        "gitlab.com/surfprace/cathal-go/exporter/parser"
        "github.com/prometheus/common/log"
        "net/http"
        "os/exec"
        "io/ioutil"
        "time"
)

// prometheus gauge metric
var jobsInQueue = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
                Name: "jobs_in_queue",
                Help: "Current number of jobs in the queue",
        },
        []string{"job_type", "slurm_group"},
)

var listenAddress = flag.String(
        "listen-address",
        ":8080",
        "The address to listen on for HTTP requests.",
)

var command string
var splitted []string

func init() {
        // register prometheus gauge metric
	prometheus.MustRegister(jobsInQueue)

        flag.StringVar(&command, "command", "squeue", "command for job queue info")
        flag.Parse()
        // split -command flag into 2 for execution
	splitted = strings.SplitN(command, " ", 2)

        // The Handler function provides a default handler to expose metrics
        // via an HTTP server. "/metrics" is the usual endpoint for that.
        log.Infof("Starting Server: %s", *listenAddress)
        http.Handle("/metrics", promhttp.Handler())
}

// Execute the squeue command and return its output
func queueData() {
        go func() {
                for {i
			// create squeue execution command
                        //cmd := exec.Command("squeue", "-a", "-r", "-h", "-o %A,%T,%r", "--states=all")
                        var cmd *exec.Cmd
                        if len(splitted) == 2 {
                                cmd = exec.Command(splitted[0], splitted[1])
                        } else {
                                cmd = exec.Command(command)
                        }

                        stdout, err := cmd.StdoutPipe()
                        if err != nil {
                                log.Fatal(err)
                        }
                        if err := cmd.Start(); err != nil {
                                log.Fatal(err)
                        }
                        out, _ := ioutil.ReadAll(stdout)
                        if err := cmd.Wait(); err != nil {
                                log.Fatal(err)
                        }

			// parse squeue output returning map
                        mp := parser.ParseQueueMetrics(out)

			// populate gauge maetric with slurm group & job type
                        for k := range mp {
                                squeue_jobs := mp[k]
                                for i := range squeue_jobs {
                                        jobsInQueue.With(prometheus.Labels{"job_type": i, "slurm_group": k}).Set(float64(squeue_jobs[i]))
                                }
                        }
			// repeat every x seconds
                        time.Sleep(5 * time.Second)
                }
        }()
}

func main() {
	// gather squeue data and create prometheus gauge metric
        queueData()
	// expose gauge metric over http endpoint 8080
        log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
