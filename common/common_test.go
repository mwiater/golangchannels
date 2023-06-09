package common_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/mattwiater/golangchannels/common"
	"github.com/mattwiater/golangchannels/structs"
	"github.com/stretchr/testify/assert"
)

func TestConsoleColumnWidth(t *testing.T) {
	columnWidth := common.ConsoleColumnWidth("12345", 30)
	testColumnWidth := 30 - len("12345")
	assert.Equal(t, testColumnWidth, columnWidth)
}

func TestSplitStringLines(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		var tests = []struct {
			testName string
			input1   string
			want     []string
		}{
			{"multiline input", "12345\n67890", []string{"12345", "67890"}},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				ans := common.SplitStringLines("12345\n67890")
				assert.EqualValues(t, fmt.Sprintln(ans), fmt.Sprintln(tt.want))
			})
		}

		tests = []struct {
			testName string
			input1   string
			want     []string
		}{
			{"single line input", "abcde", []string{"abcde"}},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				ans := common.SplitStringLines("abcde")
				assert.EqualValues(t, fmt.Sprintln(ans), fmt.Sprintln(tt.want))
			})
		}
	})
}

func TestGetAttr(t *testing.T) {
	uuid := uuid.New()
	testJobResultAttr := reflect.ValueOf(float64(1))

	job1 := structs.Job{JobNumber: 1, Id: uuid, JobName: "EmptySleepJob", Data: ""}
	jobResult1 := structs.JobResult{WorkerID: uuid, Job: job1, NumberOfJobs: 16, JobTimer: float64(1), JobMemAlloc: float32(1), JobName: "EmptySleepJob", Status: ""}

	t.Run("success", func(t *testing.T) {
		var tests = []struct {
			testName string
			input1   *structs.JobResult
			input2   string
			want     reflect.Value
		}{
			{"values are equal", &jobResult1, "JobTimer", testJobResultAttr},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				ans, _ := common.GetAttr(tt.input1, tt.input2)
				assert.EqualValues(t, fmt.Sprintln(ans), fmt.Sprintln(tt.want))
			})
		}
	})

	t.Run("calls non-existent struct field", func(t *testing.T) {
		var tests = []struct {
			testName string
			input1   *structs.JobResult
			input2   string
			want     reflect.Value
		}{
			{"error", &jobResult1, "JobTimer2", testJobResultAttr},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				_, err := common.GetAttr(tt.input1, tt.input2)
				assert.Error(t, err)
			})
		}
	})
}

func TestCalculateMemory(t *testing.T) {
	currentMemStat, _ := common.CalculateMemory()
	testCurrentMemStat := float32(12345)
	assert.IsType(t, testCurrentMemStat, currentMemStat)
}

func TestBToMb(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var tests = []struct {
			testName string
			input    float32
			want     float32
		}{
			{"less than 1", 0.955, 9.1075896e-07},
			{"greater than 1", 12, 1.1444092e-05},
			{"less than 0", -1.234, -1.1768341e-06},
		}

		for _, tc := range tests {
			got := common.BToMb(tc.input)
			assert.EqualValues(t, tc.want, got)
		}
	})
}
