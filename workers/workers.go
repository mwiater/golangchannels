package workers

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/mattwiater/golangchannels/common"
	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/jobs/emptySleepJob"
	"github.com/mattwiater/golangchannels/structs"
)

type Job = structs.Job
type JobResult = structs.JobResult
type WorkerStat = structs.WorkerStat

var jobs = make(chan Job, config.TotalJobs)
var jobResults = make(chan JobResult, config.TotalJobs)
var jobData []JobResult

var workersElapsedTime float64
var JobNumber int
var WorkerStats []WorkerStat

func WorkerResult(done chan bool) {
	for jobResult := range jobResults {
		jobResultMap := map[string]string{}
		json.Unmarshal([]byte(jobResult.Status), &jobResultMap)

		jobData = append(jobData, jobResult)

		col1 := fmt.Sprintf("    -> JOB %v/%v COMPLETED:", jobResult.Job.JobNumber, jobResult.NumberOfJobs)
		colWidth := common.ConsoleColumnWidth(col1, 35)
		config.ConsoleGreen.Printf("    -> JOB %v/%v COMPLETED: %-*s %v with Worker: %v (Ran %s in %v Seconds)\n", jobResult.Job.JobNumber, jobResult.NumberOfJobs, colWidth, "", jobResult.Job.Id, jobResult.WorkerID, jobResult.JobName, jobResult.JobTimer)
	}
	done <- true
}

func PerformJob(job Job) (string, float64) {
	myJob := emptySleepJob.Job(job)
	var result, jobTimer = myJob.EmptySleepJob()

	return result, jobTimer
}

func CreateWorkerPool(noOfWorkers int, noOfJobs int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		uuid := uuid.New()
		col1 := fmt.Sprintf("Allocating Worker #%v:", i+1)
		colWidth := common.ConsoleColumnWidth(col1, 35)
		config.ConsoleGreen.Printf("Allocating Worker #%d: %-*s %v\n", i+1, colWidth, "", uuid)

		go Worker(&wg, uuid, noOfJobs)
	}
	wg.Wait()

	if config.Debug {
		fmt.Println(jobData)
	}

	close(jobResults)
}

func AllocateJob(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		uuid := uuid.New()
		JobNumber := i + 1
		col1 := fmt.Sprintf("  Allocating Job #%d:", JobNumber)
		colWidth := common.ConsoleColumnWidth(col1, 35)
		config.ConsoleGreen.Printf("  Allocating Job #%d: %-*s %v\n", JobNumber, colWidth, "", uuid)
		job := Job{JobNumber: JobNumber, Id: uuid, JobName: "emptySleepJob"}
		jobs <- job
	}

	close(jobs)
}

func Workers(workerCount int, jobCount int, jobName string) float64 {
	// Jobs to run for each iteration
	// Need to reassign in this closure
	jobCount = jobCount
	jobs = make(chan Job, jobCount)
	jobResults = make(chan JobResult, jobCount)

	startTime := time.Now()
	numberOfJobs := jobCount

	col1 := fmt.Sprintf("Workers: %d", workerCount)
	colWidth := common.ConsoleColumnWidth(col1, 35)
	config.ConsoleWhite.Printf("Workers: %d %-*s Job Name: %s\n", workerCount, colWidth, "", jobName)

	text := fmt.Sprintf("Workers: %d %-*s Job Name: %s", workerCount, colWidth, "", jobName)
	divider := strings.Repeat("-", len(text))
	config.ConsoleWhite.Println(divider)

	col1 = fmt.Sprintf("  Workers In Use:")
	colWidth = common.ConsoleColumnWidth(col1, 35)
	config.ConsoleWhite.Printf("  Workers In Use: %-*s %v\n", colWidth, "", workerCount)

	col1 = fmt.Sprintf("  Workers Available:")
	colWidth = common.ConsoleColumnWidth(col1, 35)
	config.ConsoleWhite.Printf("  Workers Available: %-*s %v\n", colWidth, "", config.WorkersAvailable)

	col1 = fmt.Sprintf("  Workers Idle:")
	colWidth = common.ConsoleColumnWidth(col1, 35)
	config.ConsoleWhite.Printf("  Workers Idle: %-*s %v\n", colWidth, "", config.WorkersAvailable-workerCount)

	col1 = fmt.Sprintf("  Number of Jobs:")
	colWidth = common.ConsoleColumnWidth(col1, 35)
	config.ConsoleWhite.Printf("  Number of Jobs: %-*s %v\n", colWidth, "", numberOfJobs)

	config.ConsoleWhite.Println(divider)
	fmt.Println()

	go AllocateJob(numberOfJobs)
	done := make(chan bool)
	go WorkerResult(done)

	CreateWorkerPool(workerCount, numberOfJobs)
	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)

	fmt.Println()
	text = fmt.Sprintf("Total time taken: %-*s %f %s\n", colWidth, "", diff.Seconds(), "Seconds")
	divider = strings.Repeat("-", len(text))
	config.ConsoleGreen.Println(divider)

	col1 = fmt.Sprintf("Total time taken:")
	colWidth = common.ConsoleColumnWidth(col1, 35)
	config.ConsoleGreen.Printf("Total time taken: %-*s %f %s\n", colWidth, "", diff.Seconds(), "Seconds")

	fmt.Println()
	fmt.Println()

	workersElapsedTime = diff.Seconds()
	return float64(workersElapsedTime)
}

func Worker(wg *sync.WaitGroup, workerID uuid.UUID, noOfJobs int) {
	for job := range jobs {
		col1 := fmt.Sprintf("  JOB %v/%v STARTED:", job.JobNumber, config.TotalJobs)
		colWidth := common.ConsoleColumnWidth(col1, 35)
		config.ConsoleCyan.Printf("  JOB %v/%v STARTED: %-*s %v with Worker: %v\n", job.JobNumber, config.TotalJobs, colWidth, "", job.Id, workerID)
		var jobResult, jobTimer = PerformJob(job)
		output := JobResult{WorkerID: workerID, Job: job, NumberOfJobs: noOfJobs, JobTimer: jobTimer, JobName: job.JobName, Status: jobResult}
		jobResults <- output
	}
	wg.Done()
}
