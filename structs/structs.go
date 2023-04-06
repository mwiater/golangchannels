package structs

import (
	"github.com/google/uuid"
)

type Job struct {
	JobNumber int
	Id        uuid.UUID
	JobName   string
}

type JobResult struct {
	WorkerID     uuid.UUID
	Job          Job
	NumberOfJobs int
	JobTimer     float64
	JobName      string
	Status       string
}

type WorkerStat struct {
	Workers       int
	ExecutionTime float64
}

type WeaselsJobResult struct {
	Target      string
	InitialSeed string
	Generations string
	Status      string
	Elapsed     string
}

type SleepJobResult struct {
	SleepTime string
	Elapsed   string
	Status    string
}
