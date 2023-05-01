// Package errorHandler implements error handling
package errorHandler

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"

	"github.com/mattwiater/golangchannels/structs"
)

func New(errorObj error) error {
	e := getCallerInfo(2, errorObj)
	err := fmt.Errorf("Error [%d]: %s (Function: %s, File: %s#%d)", e.Code, e.Message, e.CallerName, e.CallerFile, e.CallerLine)
	// FIX
	//if config.PrettyPrintErrors {
	//	config.ConsoleRed.Println("Error:")
	fmt.Println("Error:")
	pretty(e)
	//} else {
	//	fmt.Println(err)
	//}
	return err
}

func pretty(errorObj structs.Error) {
	errorJSON, err := json.MarshalIndent(errorObj, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(errorJSON))
}

func getCallerInfo(skip int, err error) structs.Error {
	var e = structs.Error{}

	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		return e
	}

	funcName := path.Base(runtime.FuncForPC(pc).Name())
	fileName := file

	e.CallerName = funcName
	e.CallerFile = fileName
	e.CallerLine = lineNo
	e.Code = 101
	e.Message = err.Error()

	return e
}
