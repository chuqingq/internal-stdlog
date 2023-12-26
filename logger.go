package main

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	// 默认不打印日志
	f, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	logger = log.New(f, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
}

func EnableLog() {
	logger.SetOutput(log.Writer())
}
