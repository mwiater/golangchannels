package piJob_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/jobs/piJob"
	"github.com/mattwiater/golangchannels/structs"
	"github.com/stretchr/testify/assert"
)

type Job structs.Job

func TestPiJob(t *testing.T) {
	uuid := uuid.New()
	job := Job{JobNumber: 1, Id: uuid, JobName: config.JobName}
	myJob := piJob.Job(job)
	var result, jobTimer = myJob.PiJob()

	var testString string = "12345"
	var testFloat64 float64 = 100

	assert.IsType(t, testString, result)
	assert.IsType(t, testFloat64, jobTimer)
}
