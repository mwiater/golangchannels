package main

import (
	"embed"
	"fmt"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/dispatcher"
	"github.com/mattwiater/golangchannels/network"
	"github.com/mattwiater/golangchannels/workers"
	"github.com/olekukonko/tablewriter"
)

//go:embed .env
var envVarsFile embed.FS

func main() {
	// Clear Screen
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error: cmd.Run()", err)
	}

	config.EnvVarsFile = envVarsFile

	cfg, err := config.AppConfig()
	if err != nil {
		log.Fatal("Error: config.AppConfig()")
	}

	if cfg["PPROF"] == "true" {
		pprofAddress := ""
		if cfg["PPROFIP"] == "" {
			pprofAddress = network.GetOutboundIP().String()
		} else {
			pprofAddress = cfg["PPROFIP"]
		}

		pprofPort := ""
		if cfg["PPROFPORT"] == "" {
			pprofPort = "6060"
		} else {
			pprofPort = cfg["PPROFPORT"]
		}

		go func() {
			fmt.Println(http.ListenAndServe(pprofAddress+":"+pprofPort, nil))
		}()

		fmt.Println("\nPPROF Listening on: " + pprofAddress + ":" + pprofPort + "\n")
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		fmt.Println("\nShutting down.")
		os.Exit(0)
	}()

	dispatcher.Run(config.StartingWorkerCount, config.MaxWorkerCount, config.TotalJobCount)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Num Workers", "Num Jobs", "Avg Job Time", "Tot Worker Time", "+/-"})

	for i, stat := range workers.WorkerStats {
		currentStatJobElapsedAverage := stat.JobElapsedAverage
		currentStatExecutionTime := stat.ExecutionTime
		baselineExecutionTime := workers.WorkerStats[0].ExecutionTime

		workerCountString := fmt.Sprintf("%v", stat.Workers)
		jobsCountString := fmt.Sprintf("%v", config.TotalJobCount)
		jobExecutionAverage := fmt.Sprintf("%f", currentStatJobElapsedAverage)
		workerExecutionTime := fmt.Sprintf("%f", currentStatExecutionTime)

		speedIncrease := "(1x)"

		if i < len(workers.WorkerStats) && i > int(0) {
			if baselineExecutionTime > currentStatExecutionTime {
				// FASTER
				speedIncrease = fmt.Sprintf("+%vx", math.Round((baselineExecutionTime/currentStatExecutionTime)*100)/100)
			} else {
				// SLOWER
				speedIncrease = fmt.Sprintf("-%vx", math.Round((baselineExecutionTime/currentStatExecutionTime)*100)/100)
			}
		}
		table.Append([]string{workerCountString, jobsCountString, jobExecutionAverage, workerExecutionTime, speedIncrease})
	}

	fmt.Println("\nSummary:")
	table.Render()
}
