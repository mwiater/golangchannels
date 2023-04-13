package dispatcher_test

import (
	"flag"
	"fmt"
	"strconv"
	"testing"

	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/dispatcher"
)

var startingWorkerCountString = flag.String("startingWorkerCount", "", "startingWorkerCount")
var maxWorkerCountString = flag.String("maxWorkerCount", "", "maxWorkerCount")
var jobCountString = flag.String("jobCount", "", "jobCount")

func BenchmarkRun(b *testing.B) {
	startingWorkerCount, _ := strconv.Atoi(*startingWorkerCountString)
	maxWorkerCount, _ := strconv.Atoi(*maxWorkerCountString)
	jobCount, _ := strconv.Atoi(*jobCountString)
	jobName := "PiJob"

	for i := startingWorkerCount; i <= maxWorkerCount; i++ {
		b.Run(fmt.Sprintf("Worker Count: %d", startingWorkerCount), func(b *testing.B) {
			config.ConsoleGreen.Printf("\n  SETUP:\n")
			config.ConsoleGreen.Printf("  #=> Starting Worker Count: %*d\n", 7, startingWorkerCount)
			config.ConsoleGreen.Printf("  #=> Max Worker Count: %*d\n", 12, maxWorkerCount)
			config.ConsoleGreen.Printf("  #=> Job Count: %*d\n", 20, jobCount)
			fmt.Println()
			dispatcher.Run(jobName, startingWorkerCount, maxWorkerCount, jobCount)
			startingWorkerCount++
		})
	}
}
