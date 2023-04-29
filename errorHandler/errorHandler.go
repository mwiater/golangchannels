// Package errorHandler implements error handling
package errorHandler

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/mattwiater/golangchannels/structs"
)

func New(errorObj error, prettyPrint bool) structs.Error {
	e := getCallerInfo(1, errorObj)

	fmt.Println("e", e)

	if prettyPrint {
		pretty(e)
	} else {
		fmt.Println(e.Message)
	}

	return e
}

func pretty(errorObj structs.Error) {
	errorJSON, err := json.MarshalIndent(errorObj, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(errorJSON))
	os.Exit(1)
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
