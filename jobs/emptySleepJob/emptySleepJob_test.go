package emptySleepJob_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mattwiater/golangchannels/jobs/emptySleepJob"
	"github.com/mattwiater/golangchannels/structs"
	"github.com/stretchr/testify/assert"
)

type Job structs.Job

func TestEmptySleepJob(t *testing.T) {
	uuid := uuid.New()
	job := Job{JobNumber: 1, Id: uuid, JobName: "emptySleepJob"}
	myJob := emptySleepJob.Job(job)
	var result, jobTimer = myJob.EmptySleepJob()

	var testString string = "12345"
	var testFloat64 float64 = 100

	assert.IsType(t, testString, result)
	assert.IsType(t, testFloat64, jobTimer)
}
