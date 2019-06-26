package main

import (
	"flag"
	"log"

	"github.com/priteshgudge/goworkerpool/workerpool/pool"
	work "github.com/priteshgudge/goworkerpool/workerpool/work"
)

// note, that variables are pointers
var workerCount = flag.Int("workerCount", 5, "Number of Workers")
var jobCount = flag.Int("numberofJobs", 100, "Total Number of Jobs")

func init() {
	// example with short version for long flag
}

func main() {
	log.Println("starting application...")
	collector := pool.StartDispatcher(*workerCount) // start up worker pool

	for i, job := range work.CreateJobs(*jobCount) {
		collector.Work <- pool.Work{Job: job, ID: i}
	}
}
