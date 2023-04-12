package dispatcher_test

import (
	"flag"
	"fmt"
	"strconv"
	"testing"

	"github.com/mattwiater/golangchannels/dispatcher"
)

var startingWorkerCountString = flag.String("startingWorkerCount", "", "startingWorkerCount")
var maxWorkerCountString = flag.String("maxWorkerCount", "", "maxWorkerCount")
var jobCountString = flag.String("jobCount", "", "jobCount")

func BenchmarkRun(b *testing.B) {
	startingWorkerCount, _ := strconv.Atoi(*startingWorkerCountString)
	maxWorkerCount, _ := strconv.Atoi(*maxWorkerCountString)
	jobCount, _ := strconv.Atoi(*jobCountString)

	for i := startingWorkerCount; i <= maxWorkerCount; i++ {
		b.Run(fmt.Sprintf("Worker Count: %d", startingWorkerCount), func(b *testing.B) {
			fmt.Println("startingWorkerCount:", startingWorkerCount, "maxWorkerCount:", maxWorkerCount, "jobCount:", jobCount)
			dispatcher.Run(startingWorkerCount, maxWorkerCount, jobCount)
			startingWorkerCount++
		})
	}
}
