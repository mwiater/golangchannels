package config

import (
	"embed"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/mattwiater/golangchannels/common"
)

var WorkersAvailable = runtime.NumCPU()
var TotalJobs = 16
var StartingWorkerCount = 1
var EndingWorkerCount = WorkersAvailable

var EmptySleepJobSleepTimeMs = 1000

var Debug = false
var ConsoleGreen = color.New(color.FgGreen)
var ConsoleCyan = color.New(color.FgCyan)
var ConsoleWhite = color.New(color.FgWhite)

var EnvVarsFile embed.FS

// AppConfig returns a new decoded Config struct
func AppConfig() (map[string]string, error) {
	envVars, _ := EnvVarsFile.ReadFile(".env")

	lines := common.SplitStringLines(string(envVars))

	var envs = make(map[string]string)
	for _, line := range lines {
		keyValuePair := strings.Split(line, "=")
		envs[keyValuePair[0]] = keyValuePair[1]
	}
	return envs, nil
}
