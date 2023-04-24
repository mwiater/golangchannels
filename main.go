package main

import (
	"embed"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"

	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/dispatcher"
	"github.com/mattwiater/golangchannels/workers"
	"github.com/olekukonko/tablewriter"
)

//go:embed .env
var envVarsFile embed.FS

func main() {
	config.EnvVarsFile = envVarsFile

	_, err := config.AppConfig()
	if err != nil {
		log.Fatal("Error: config.AppConfig()")
	}

	jobName := config.JobName

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		fmt.Println("\nShutting down.")
		os.Exit(0)
	}()

	dispatcher.Run(jobName, config.StartingWorkerCount, config.MaxWorkerCount, config.TotalJobCount)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Workers", "Jobs", "Avg Job Time", "Total Worker Time", "Avg Mem Use", "+/-"})

	for i, stat := range workers.WorkerStats {
		currentStatJobElapsedAverage := stat.JobElapsedAverage
		currentStatExecutionTime := stat.ExecutionTime
		baselineExecutionTime := workers.WorkerStats[0].ExecutionTime
		currentStatMemAllocAverage := float64(stat.MemAllocAverage)

		workerCountString := fmt.Sprintf("%v", stat.Workers)
		jobsCountString := fmt.Sprintf("%v", config.TotalJobCount)
		jobExecutionAverage := fmt.Sprintf("%.2fs", currentStatJobElapsedAverage)
		memAllocAverage := fmt.Sprintf("%.3fMb", math.Round(currentStatMemAllocAverage*1000)/1000)
		workerExecutionTime := fmt.Sprintf("%.2fs", currentStatExecutionTime)

		speedIncrease := "(1x)*"

		if i < len(workers.WorkerStats) && i > int(0) {
			if baselineExecutionTime > currentStatExecutionTime {
				// FASTER
				speedIncrease = fmt.Sprintf("+%vx", math.Round((baselineExecutionTime/currentStatExecutionTime)*100)/100)
			} else {
				// SLOWER
				speedIncrease = fmt.Sprintf("-%vx", math.Round((baselineExecutionTime/currentStatExecutionTime)*100)/100)
			}
		}
		table.Append([]string{workerCountString, jobsCountString, jobExecutionAverage, workerExecutionTime, memAllocAverage, speedIncrease})
	}

	fmt.Println()
	fmt.Println("\nSummary Results:", jobName)
	table.Render()

	fmt.Println()
	fmt.Println("* Baseline: All subsequent +/- tests are compared to this.")
	fmt.Println()
}
