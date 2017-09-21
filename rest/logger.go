package rest

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}
}

//Log log a message in console
func Log(v ...interface{}) {
	logger.Println(v...)
}

//Logf log a message in console
func Logf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}
