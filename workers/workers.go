// Package workers is used to create workers that handle jobs
package workers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/mattwiater/golangchannels/common"
	"github.com/mattwiater/golangchannels/config"

	//"github.com/mattwiater/golangchannels/jobs/emptySleepJob"
	PiJob "github.com/mattwiater/golangchannels/jobs/piJob"
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

func WorkerResult(done chan []structs.JobResult) {
	for jobResult := range jobResultsChannel {
		jobResultMap := map[string]string{}
		json.Unmarshal([]byte(jobResult.Status), &jobResultMap)

		JobResults = append(JobResults, jobResult)

		if config.Debug {
			col1 := fmt.Sprintf("    -> JOB %v/%v COMPLETED:", jobResult.Job.JobNumber, jobResult.NumberOfJobs)
			colWidth := common.ConsoleColumnWidth(col1, 35)
			config.ConsoleGreen.Printf("    -> JOB %v/%v COMPLETED: %-*s %v with Worker: %v (Ran %s in %v Seconds)\n", jobResult.Job.JobNumber, jobResult.NumberOfJobs, colWidth, "", jobResult.Job.Id, jobResult.WorkerID, jobResult.JobName, jobResult.JobTimer)
		}
	}

	done <- JobResults
}

func PerformJob(job Job) (string, float64) {
	myJob := PiJob.Job(job)
	var result, jobTimer = myJob.PiJob()

	// myJob := emptySleepJob.Job(job)
	// var result, jobTimer = myJob.EmptySleepJob()

	return result, jobTimer
}

func CreateWorkerPool(noOfWorkers int, noOfJobs int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		uuid := uuid.New()

		if config.Debug {
			col1 := fmt.Sprintf("Allocating Worker #%v:", i+1)
			colWidth := common.ConsoleColumnWidth(col1, 35)
			config.ConsoleGreen.Printf("Allocating Worker #%d: %-*s %v\n", i+1, colWidth, "", uuid)
		}

		go Worker(&wg, uuid, noOfJobs)
	}
	wg.Wait()

	close(jobResultsChannel)
}

func AllocateJob(jobName string, noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		uuid := uuid.New()
		JobNumber := i + 1

		if config.Debug {
			col1 := fmt.Sprintf("  Allocating Job #%d:", JobNumber)
			colWidth := common.ConsoleColumnWidth(col1, 35)
			config.ConsoleGreen.Printf("  Allocating Job #%d: %-*s %v\n", JobNumber, colWidth, "", uuid)
		}

		job := Job{JobNumber: JobNumber, Id: uuid, JobName: jobName}
		jobs <- job
	}

	close(jobs)
}

func Workers(workerCount int, jobCount int, jobName string) (float64, float64) {
	// Jobs to run for each iteration
	// Need to reassign in this closure
	jobCount = jobCount
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

	jobElapsedSum := 0.0
	for _, allJobResult := range allJobResults[(len(allJobResults) - numberOfJobs):] { // iterate over desired rows
		jobTime := getAttr(&allJobResult, "JobTimer")
		jobTimeFloat := jobTime.Interface().(float64)
		jobElapsedSum += (jobTimeFloat)

	}

	jobElapsedAvg := (float64(jobElapsedSum)) / (float64(numberOfJobs))
	endTime := time.Now()
	diff := endTime.Sub(startTime)

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
	return float64(workersElapsedTime), jobElapsedAvg
}

func getAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}
	return curField
}

func Worker(wg *sync.WaitGroup, workerID uuid.UUID, noOfJobs int) {
	for job := range jobs {
		if config.Debug {
			col1 := fmt.Sprintf("  JOB %v/%v STARTED:", job.JobNumber, config.TotalJobCount)
			colWidth := common.ConsoleColumnWidth(col1, 35)
			config.ConsoleCyan.Printf("  JOB %v/%v STARTED: %-*s %v with Worker: %v\n", job.JobNumber, config.TotalJobCount, colWidth, "", job.Id, workerID)
		}

		var jobResult, jobTimer = PerformJob(job)
		output := JobResult{WorkerID: workerID, Job: job, NumberOfJobs: noOfJobs, JobTimer: jobTimer, JobName: job.JobName, Status: jobResult}
		jobResultsChannel <- output
	}
	wg.Done()
}
