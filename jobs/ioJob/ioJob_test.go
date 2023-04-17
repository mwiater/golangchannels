package ioJob_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/jobs/ioJob"
	"github.com/mattwiater/golangchannels/structs"
	"github.com/stretchr/testify/assert"
)

type Job structs.Job

func TestIoJob(t *testing.T) {
	uuid := uuid.New()
	job := Job{JobNumber: 1, Id: uuid, JobName: config.JobName}
	myJob := ioJob.Job(job)
	var result, jobTimer = myJob.IoJob()

	var testString string = "12345"
	var testFloat64 float64 = 100.001

	assert.IsType(t, testString, result)
	assert.IsType(t, testFloat64, jobTimer)
}
