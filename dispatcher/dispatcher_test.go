package dispatcher_test

import (
	"flag"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/dispatcher"
)

var jobName = flag.String("jobName", "", "jobName")
var startingWorkerCountString = flag.String("startingWorkerCount", "", "startingWorkerCount")
var maxWorkerCountString = flag.String("maxWorkerCount", "", "maxWorkerCount")
var jobCountString = flag.String("jobCount", "", "jobCount")

func BenchmarkRun(b *testing.B) {
	if *jobName == "" {
		b.Error("\nMissing argument: -jobName\nFormat must include all arguments, e,g: -jobName=PiJob -startingWorkerCount=1 -maxWorkerCount=1 -jobCount=64\n")
		b.FailNow()
	}

	if *startingWorkerCountString == "" {
		b.Error("\nMissing argument: -startingWorkerCountString\nFormat must include all arguments, e,g: -jobName=PiJob -startingWorkerCount=1 -maxWorkerCount=1 -jobCount=64\n")
		b.FailNow()
	}

	if *maxWorkerCountString == "" {
		b.Error("\nMissing argument: -maxWorkerCountString\nFormat must include all arguments, e,g: -jobName=PiJob -startingWorkerCount=1 -maxWorkerCount=1 -jobCount=64\n")
		b.FailNow()
	}

	if *jobCountString == "" {
		b.Error("\nMissing argument: -jobCountString\nFormat must include all arguments, e,g: -jobName=PiJob -startingWorkerCount=1 -maxWorkerCount=1 -jobCount=64\n")
		b.FailNow()
	}

	startingWorkerCount, _ := strconv.Atoi(*startingWorkerCountString)
	maxWorkerCount, _ := strconv.Atoi(*maxWorkerCountString)
	jobCount, _ := strconv.Atoi(*jobCountString)

	for i := startingWorkerCount; i <= maxWorkerCount; i++ {
		b.Run(fmt.Sprintf("Worker Count: %d", startingWorkerCount), func(b *testing.B) {

			benchmarkStartTime := time.Now()
			config.ConsoleGreen.Printf("\n  BENCHMARK SETUP:\n")
			config.ConsoleGreen.Printf("  #=> Job: %*s\n", 7, *jobName)
			config.ConsoleGreen.Printf("  #=> Starting Worker Count: %*d\n", 7, startingWorkerCount)
			config.ConsoleGreen.Printf("  #=> Max Worker Count: %*d\n", 12, maxWorkerCount)
			config.ConsoleGreen.Printf("  #=> Job Count: %*d\n", 20, jobCount)
			dispatcher.Run(*jobName, startingWorkerCount, maxWorkerCount, jobCount)
			benchmarkEndTime := time.Now()
			benchmarkElapsed := benchmarkEndTime.Sub(benchmarkStartTime)
			config.ConsoleCyan.Printf("    #=> Benchmark Elapsed: %*s\n", 21, benchmarkElapsed)
			fmt.Println()
			startingWorkerCount++
		})
	}
}
