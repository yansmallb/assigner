package main

import (
	"github.com/yansmallb/assigner/cli"
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
)

func main() {
	logFilename := "/tmp/assigner.log"
	logFile, _ := os.OpenFile(logFilename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer logFile.Close()

	writers := []io.Writer{
		logFile,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)

	log.SetOutput(fileAndStdoutWriter)
	log.SetLevel(log.DebugLevel)

	log.Infoln("main.main():Start Assigner Main")
	cli.Run()
}
