// Package common implements shared application functions
package common

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/mattwiater/golangchannels/structs"
)

func GetCallerInfo(skip int) structs.Error {
	var e = structs.Error{}

	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		return e
	}

	funcName := path.Base(runtime.FuncForPC(pc).Name())
	fileName := file

	e.CallerName = fmt.Sprintf("%s", funcName)
	e.CallerFile = fmt.Sprintf("%s", fileName)
	e.CallerLine = fmt.Sprintf("%d", lineNo)

	return e
}

// Get console column width of submitted string
func ConsoleColumnWidth(text string, finalColWidth int) int {
	return finalColWidth - len(text)
}

// Return slice with each line of a multi-line string, splitting on '\n'
func SplitStringLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

// GetAttr iterates through an interface and returns the value of the requested field.
func GetAttr(obj interface{}, fieldName string) (reflect.Value, error) {
	pointToStruct := reflect.ValueOf(obj)
	curStruct := pointToStruct.Elem()
	curField := curStruct.FieldByName(fieldName)
	if !curField.IsValid() {
		emptyValue := reflect.ValueOf([]interface{}{nil}).Index(0)
		return emptyValue, fmt.Errorf("Field not found: %v", fieldName)
	}
	return curField, nil
}

// CalculateMemory returns the current process memory usage.
func CalculateMemory() (float32, error) {
	process_id := os.Getpid()
	f, err := os.Open(fmt.Sprintf("/proc/%d/smaps", process_id))
	if err != nil {
		return 0, err
	}
	defer f.Close()

	memoryBytes := uint64(0)
	pfx := []byte("Pss:")
	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Bytes()
		if bytes.HasPrefix(line, pfx) {
			var size uint64
			_, err := fmt.Sscanf(string(line[4:]), "%d", &size)
			if err != nil {
				return 0, err
			}
			memoryBytes += size
		}
	}
	if err := r.Err(); err != nil {
		return 0, err
	}

	memoryMB := BToMb(float32(memoryBytes))
	return float32(memoryMB), nil
}

// BToMb converts bytes to megabytes.
func BToMb(b float32) float32 {
	return b / 1024 / 1024
}
