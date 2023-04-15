// Package ioJob is used to simuilate an i/0 intesive job
package ioJob

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/structs"
)

type Job structs.Job

func (job Job) IoJob() (string, float64) {
	jobStartTime := time.Now()
	iterations := 11
	for n := 0; n < iterations; n++ {
		f, err := os.Create("/tmp/test.txt")
		if err != nil {
			panic(err)
		}
		for i := 0; i < 100000; i++ {
			f.WriteString("some text!\n")
		}
		f.Close()
	}

	jobEndTime := time.Now()
	jobElapsed := jobEndTime.Sub(jobStartTime)

	jobResult := structs.SleepJobResult{}
	jobResult.SleepTime = time.Duration(config.EmptySleepJobSleepTimeMs).String()
	jobResult.Elapsed = jobElapsed.String()
	jobResult.Status = strconv.FormatBool(true)

	jobResultString, err := json.Marshal(jobResult)
	if err != nil {
		fmt.Println(err)
	}

	return string(jobResultString), jobElapsed.Seconds()
}
