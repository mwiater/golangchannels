package dispatcher

import (
	"fmt"
	"strings"

	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/workers"
)

func Run() {

	for i := config.StartingWorkerCount; i <= config.EndingWorkerCount; i++ {
		currentWorkers := i

		text := fmt.Sprintf("|  Spawning workers for test %v of %v  |", currentWorkers, config.EndingWorkerCount)
		divider := strings.Repeat("-", len(text))

		fmt.Println()
		config.ConsoleCyan.Println(divider)
		config.ConsoleCyan.Printf("|  Spawning workers for test %v of %v  |\n", currentWorkers, config.EndingWorkerCount)
		config.ConsoleCyan.Println(divider)
		fmt.Println()

		elapsed := workers.Workers(currentWorkers, config.TotalJobs, "emptySleepJob")
		workers.WorkerStats = append(workers.WorkerStats, workers.WorkerStat{Workers: currentWorkers, ExecutionTime: elapsed})
	}
}
