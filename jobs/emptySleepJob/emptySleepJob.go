// Package emptySleepJob is used to mock a consistent job by sleeping for a period of time
package emptySleepJob

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/mattwiater/golangchannels/config"
	"github.com/mattwiater/golangchannels/structs"
)

type Job structs.Job

func (job Job) EmptySleepJob() (string, float64) {
	jobStartTime := time.Now()

	time.Sleep(time.Duration(config.EmptySleepJobSleepTimeMs) * time.Millisecond)

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
