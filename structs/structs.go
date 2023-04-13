package structs

import (
	"github.com/google/uuid"
)

type Job struct {
	JobNumber int
	Id        uuid.UUID
	JobName   string
	Data      string
}

type JobResult struct {
	WorkerID     uuid.UUID
	Job          Job
	NumberOfJobs int
	JobTimer     float64
	JobName      string
	Status       string
	Data         string
}

type WorkerStat struct {
	Workers           int
	JobName           string
	ExecutionTime     float64
	JobElapsedAverage float64
	Data              string
}

type SleepJobResult struct {
	SleepTime string
	Elapsed   string
	Status    string
	Data      string
}
