// Package config implements .env file application configuration
package config

import (
	"embed"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/mattwiater/golangchannels/common"
)

var WorkersAvailable = runtime.NumCPU()
var EmptySleepJobSleepTimeMs = 1000

var ConsoleGreen = color.New(color.FgGreen)
var ConsoleCyan = color.New(color.FgCyan)
var ConsoleWhite = color.New(color.FgWhite)
var ConsoleRed = color.New(color.FgRed)

var EnvVarsFile embed.FS

var Debug bool
var PrettyPrintErrors bool
var JobName string
var StartingWorkerCount int
var MaxWorkerCount int
var TotalJobCount int

// AppConfig returns a new decoded Config map from .env file variables or sets from defaults
func AppConfig() (map[string]string, error) {
	envVars, _ := EnvVarsFile.ReadFile(".env")
	lines := common.SplitStringLines(string(envVars))
	var envs = make(map[string]string)
	for _, line := range lines {
		keyValuePair := strings.Split(line, "=")
		envs[keyValuePair[0]] = keyValuePair[1]

		if keyValuePair[0] == "DEBUG" {
			if keyValuePair[1] == "" {
				Debug = false
			} else {
				Debug, _ = strconv.ParseBool(keyValuePair[1])
			}
		}

		if keyValuePair[0] == "PRETTYPRINTERRORS" {
			if keyValuePair[1] == "" {
				PrettyPrintErrors = false
			} else {
				PrettyPrintErrors, _ = strconv.ParseBool(keyValuePair[1])
			}
		}

		if keyValuePair[0] == "JOBNAME" {
			if keyValuePair[1] == "" {
				JobName = "EmptySleepJob"
			} else {
				JobName = keyValuePair[1]
			}
		}

		if keyValuePair[0] == "STARTINGWORKERCOUNT" {
			if keyValuePair[1] == "" {
				StartingWorkerCount = 1
			} else {
				StartingWorkerCount, _ = strconv.Atoi(keyValuePair[1])
			}
		}

		if keyValuePair[0] == "MAXWORKERCOUNT" {
			if keyValuePair[1] == "" {
				MaxWorkerCount = WorkersAvailable
			} else {
				MaxWorkerCount, _ = strconv.Atoi(keyValuePair[1])
			}
		}

		if keyValuePair[0] == "TOTALJOBCOUNT" {
			if keyValuePair[1] == "" {
				TotalJobCount = WorkersAvailable * 2
			} else {
				TotalJobCount, _ = strconv.Atoi(keyValuePair[1])
			}
		}
	}
	return envs, nil
}
