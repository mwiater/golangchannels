// Package workers is used to create workers that handle jobs
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
	"github.com/mattwiater/golangchannels/jobs/ioJob"
	"github.com/mattwiater/golangchannels/jobs/piJob"
	"github.com/mattwiater/golangchannels/structs"
)

type Job = structs.Job
type JobResult = structs.JobResult
type WorkerStat = structs.WorkerStat

var jobs = make(chan Job, config.TotalJobCount)
var jobResultsChannel = make(chan JobResult, config.TotalJobCount)
var JobResults []JobResult

var workersElapsedTime float64
var JobNumber int
var WorkerStats []WorkerStat

// Workers create worker pools, set up channels, and allocate and perform jobs
// Each worker returns timing inforation back to the dispacher: total worker elapsed time for processing all jobs, average job processing time
func Workers(jobName string, workerCount int, jobCount int) (float64, float64, float32) {
	// Jobs to run for each iteration
	// Need to reassign in this closure
	jobCount = jobCount //nolint
	jobs = make(chan Job, jobCount)
	jobResultsChannel = make(chan JobResult, jobCount)

	startTime := time.Now()
	numberOfJobs := jobCount

	if config.Debug {
		col1 := fmt.Sprintf("Workers: %d", workerCount)
		colWidth := common.ConsoleColumnWidth(col1, 35)
		config.ConsoleWhite.Printf("Workers: %d %-*s Job Name: %s\n", workerCount, colWidth, "", jobName)

		text := fmt.Sprintf("Workers: %d %-*s Job Name: %s", workerCount, colWidth, "", jobName)
		divider := strings.Repeat("-", len(text))
		config.ConsoleWhite.Println(divider)

		col1 = fmt.Sprintf("  Workers In Use:%v", "")
		colWidth = common.ConsoleColumnWidth(col1, 35)
		config.ConsoleWhite.Printf("  Workers In Use: %-*s %v\n", colWidth, "", workerCount)

		col1 = fmt.Sprintf("  Workers Available:%v", "")
		colWidth = common.ConsoleColumnWidth(col1, 35)
		config.ConsoleWhite.Printf("  Workers Available: %-*s %v\n", colWidth, "", config.WorkersAvailable)

		col1 = fmt.Sprintf("  Workers Idle:%v", "")
		colWidth = common.ConsoleColumnWidth(col1, 35)
		config.ConsoleWhite.Printf("  Workers Idle: %-*s %v\n", colWidth, "", config.WorkersAvailable-workerCount)

		col1 = fmt.Sprintf("  Number of Jobs:%v", "")
		colWidth = common.ConsoleColumnWidth(col1, 35)
		config.ConsoleWhite.Printf("  Number of Jobs: %-*s %v\n", colWidth, "", numberOfJobs)

		config.ConsoleWhite.Println(divider)
		fmt.Println()
	}

	go AllocateJob(jobName, numberOfJobs)
	jobResults := make(chan []structs.JobResult)
	go WorkerResult(jobResults)

	CreateWorkerPool(workerCount, numberOfJobs)
	allJobResults := <-jobResults

	memAllocSum := float32(0)
	jobElapsedSum := 0.0
	for _, allJobResult := range allJobResults[(len(allJobResults) - numberOfJobs):] {
		jobTime := common.GetAttr(&allJobResult, "JobTimer")
		jobTimeFloat := jobTime.Interface().(float64)
		jobElapsedSum += (jobTimeFloat)

		memAlloc := common.GetAttr(&allJobResult, "JobMemAlloc")
		memAllocFloat := memAlloc.Interface().(float32)
		memAllocSum += memAllocFloat
	}
	endTime := time.Now()
	diff := endTime.Sub(startTime)

	jobElapsedAvg := (float64(jobElapsedSum)) / (float64(numberOfJobs))
	memAllocAvg := float32(memAllocSum / float32(numberOfJobs))

	if config.Debug {
		fmt.Println()
		colWidth := common.ConsoleColumnWidth("", 35)
		text := fmt.Sprintf("Total time taken: %-*s %f %s\n", colWidth, "", diff.Seconds(), "Seconds")
		divider := strings.Repeat("-", len(text))
		config.ConsoleGreen.Println(divider)

		col1 := fmt.Sprintf("Total time taken:%v", "")
		colWidth = common.ConsoleColumnWidth(col1, 35)
		config.ConsoleGreen.Printf("Total time taken: %-*s %f %s\n\n\n", colWidth, "", diff.Seconds(), "Seconds")
	}

	workersElapsedTime = diff.Seconds()
	return float64(workersElapsedTime), jobElapsedAvg, float32(memAllocAvg)
}

// CreateWorkerPool creates the desireed number or workers, assigns UUIDs, and create a worker pool waitgroup to synchronize workers.
func CreateWorkerPool(noOfWorkers int, noOfJobs int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		uuid := uuid.New()
		uuidSegments := strings.Split(uuid.String(), "-")
		uuidTrimmed := strings.Join(uuidSegments[:2], "-")

		if config.Debug {
			col1 := fmt.Sprintf("Allocating Worker #%v:", i+1)
			colWidth := common.ConsoleColumnWidth(col1, 35)
			config.ConsoleGreen.Printf("Allocating Worker #%d: %-*s %v\n", i+1, colWidth, "", uuidTrimmed)
		}

		go Worker(&wg, uuidTrimmed, noOfJobs)
	}
	wg.Wait()

	close(jobResultsChannel)
}

// Worker recieves jobs from the jobs channel and performs the requested jobs as they arrive.
// The results of each job are then sent through the jobResultsChannel channel as a JobResult struct
func Worker(wg *sync.WaitGroup, workerID string, noOfJobs int) {
	for job := range jobs {
		if config.Debug {
			col1 := fmt.Sprintf("  JOB %v/%v STARTED:", job.JobNumber, config.TotalJobCount)
			colWidth := common.ConsoleColumnWidth(col1, 35)
			config.ConsoleCyan.Printf("  JOB %v/%v STARTED: %-*s %v with Worker: %v\n", job.JobNumber, config.TotalJobCount, colWidth, "", job.Id, workerID)
		}
		var jobResultOutput, jobTimer = PerformJob(job.JobName, job)
		currentMemStat, _ := common.CalculateMemory()

		jobResult := JobResult{WorkerID: workerID, Job: job, NumberOfJobs: noOfJobs, JobTimer: jobTimer, JobMemAlloc: currentMemStat, JobName: job.JobName, Status: jobResultOutput}
		jobResultsChannel <- jobResult
	}
	wg.Done()
}

// AllocateJob creates the initial Job struct objects with metadata: JobNumber (int), Id (uuid), JobName (string)
// These jobs are then sent through the job channel.
func AllocateJob(jobName string, noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		uuid := uuid.New()
		uuidSegments := strings.Split(uuid.String(), "-")
		uuidTrimmed := strings.Join(uuidSegments[:2], "-")

		JobNumber := i + 1

		if config.Debug {
			col1 := fmt.Sprintf("  Allocating Job #%d:", JobNumber)
			colWidth := common.ConsoleColumnWidth(col1, 35)
			config.ConsoleGreen.Printf("  Allocating Job #%d: %-*s %v\n", JobNumber, colWidth, "", uuidTrimmed)
		}

		job := Job{JobNumber: JobNumber, Id: uuidTrimmed, JobName: jobName}
		jobs <- job
	}

	close(jobs)
}

// PerformJob send the job into the `jobRouter` function so that the jobName (string) can be executed as a function.
// It receieves job result data and job timing information and sends it back to the Worker
func PerformJob(jobName string, job Job) (string, float64) {
	var result string
	var jobTimer float64
	result, jobTimer = jobRouter(jobName, job)
	return result, jobTimer
}

// WorkerResult consumes the jobResultsChannel and appends each jobResult to the JobResults slice before sending it through the workerResultsChannel channel
func WorkerResult(workerResultsChannel chan []structs.JobResult) {
	for jobResult := range jobResultsChannel {
		jobResultMap := map[string]string{}
		err := json.Unmarshal([]byte(jobResult.Status), &jobResultMap)
		if err != nil {
			panic("Could not Unmarshal object")
		}

		JobResults = append(JobResults, jobResult)

		if config.Debug {
			col1 := fmt.Sprintf("    -> JOB %v/%v COMPLETED:", jobResult.Job.JobNumber, jobResult.NumberOfJobs)
			colWidth := common.ConsoleColumnWidth(col1, 35)
			config.ConsoleGreen.Printf("    -> JOB %v/%v COMPLETED: %-*s %v with Worker: %v (Ran %s in %.3f Seconds / %.3fMB)\n", jobResult.Job.JobNumber, jobResult.NumberOfJobs, colWidth, "", jobResult.Job.Id, jobResult.WorkerID, jobResult.JobName, jobResult.JobTimer, jobResult.JobMemAlloc)
		}
	}
	workerResultsChannel <- JobResults
}

// jobRouter takes a jobName (string) and then executes that named function.
// It receieves job result data and job timing information and sends it back to PerformJob
func jobRouter(jobName string, job structs.Job) (string, float64) {
	switch jobName {
	case "EmptySleepJob":
		myJob := emptySleepJob.Job(job)
		result, jobTimer := myJob.EmptySleepJob()
		return result, jobTimer
	case "PiJob":
		myJob := piJob.Job(job)
		result, jobTimer := myJob.PiJob()
		return result, jobTimer
	case "IoJob":
		myJob := ioJob.Job(job)
		result, jobTimer := myJob.IoJob()
		return result, jobTimer
	default:
		panic("Unknown function name: " + jobName)
	}
}
