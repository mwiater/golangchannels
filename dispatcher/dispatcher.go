package dispatcher

import (
	"fmt"
	"strings"

	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/workers"
)

func Run() {
	testCount := 1
	for i := config.StartingWorkerCount; i <= config.MaxWorkerCount; i++ {
		currentWorkers := i

		text := fmt.Sprintf("|  Spawning workers for test %v of %v  |", testCount, (config.MaxWorkerCount - config.StartingWorkerCount + 1))
		divider := strings.Repeat("-", len(text))

		fmt.Println()
		config.ConsoleCyan.Println(divider)
		config.ConsoleCyan.Printf("|  Spawning workers for test %v of %v  |\n", testCount, (config.MaxWorkerCount - config.StartingWorkerCount + 1))
		config.ConsoleCyan.Println(divider)
		fmt.Println()

		elapsed := workers.Workers(currentWorkers, config.TotalJobCount, "emptySleepJob")
		workers.WorkerStats = append(workers.WorkerStats, workers.WorkerStat{Workers: currentWorkers, ExecutionTime: elapsed})

		testCount++
	}
}
