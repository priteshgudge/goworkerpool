package main

import (
	"testing"

	"github.com/priteshgudge/goworkerpool/workerpool/pool"
	work "github.com/priteshgudge/goworkerpool/workerpool/work"
)

func BenchmarkConcurrent(b *testing.B) {
	collector := pool.StartDispatcher(5) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, job := range work.CreateJobs(20) {
			collector.Work <- pool.Work{Job: job, ID: i}
		}
	}
}

func BenchmarkNonconcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, job := range work.CreateJobs(20) {
			work.DoWork(job, 1)
		}
	}
}
