package utils

import (
	"log"
	"os"
)

var (
	info  = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	warn  = log.New(os.Stdout, "[WARN] ", log.LstdFlags)
	error = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
)

func InitLogger() {}

func LogInfo(msg string) {
	info.Println(msg)
}
func LogInfof(format string, v ...interface{}) {
	info.Printf(format+"\n", v...)
}
func LogWarnf(format string, v ...interface{}) {
	warn.Printf(format+"\n", v...)
}
func LogErrorf(format string, v ...interface{}) {
	error.Printf(format+"\n", v...)
}
