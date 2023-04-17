// Package structs is used to define all of the Structs used in the application.
package structs

type Job struct {
	JobNumber int
	Id        string
	JobName   string
	Data      string
}

type JobResult struct {
	WorkerID     string
	Job          Job
	NumberOfJobs int
	JobTimer     float64
	JobMemAlloc  float32
	JobName      string
	Status       string
	Data         string
}

type WorkerStat struct {
	Workers           int
	JobName           string
	ExecutionTime     float64
	JobElapsedAverage float64
	MemAllocAverage   float32
	Data              string
}

type SleepJobResult struct {
	SleepTime string
	Elapsed   string
	Status    string
	Data      string
}
