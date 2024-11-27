package log

import (
	"factory_management_go/app/util"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type StdLogger struct {
	logger       *log.Logger
	loggingLevel LoggingLevel
	pid          int
}

var Logger StdLogger

type LoggingLevel string

const (
	ERROR LoggingLevel = "ERROR"
	WARN  LoggingLevel = "WARN"
	INFO  LoggingLevel = "INFO"
	DEBUG LoggingLevel = "DEBUG"
)

var loggingLevels = map[LoggingLevel]uint8{ERROR: 0, WARN: 1, INFO: 2, DEBUG: 3}

//DEBUG
//INFO
//WARN
//ERROR

func Initialise(fileName string) error {
	loggingLevelInterface, err := util.GetProperty(util.LoggingLevel)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	writer := io.MultiWriter(file, os.Stdout)
	loggingLevel := LoggingLevel(loggingLevelInterface.(string))
	if _, exists := loggingLevels[loggingLevel]; !exists {
		log.Printf("unable to get logging level, defaulting to INFO")
		loggingLevel = INFO
	}
	Logger = StdLogger{logger: log.New(writer, "", 0), loggingLevel: loggingLevel, pid: os.Getpid()}
	return nil
}

func (logger *StdLogger) Error(message string, fileName string) {
	logger.doLog(message, fileName, ERROR)
}
func (logger *StdLogger) Warn(message string, fileName string) {
	logger.doLog(message, fileName, WARN)
}
func (logger *StdLogger) Info(message string, fileName string) {
	logger.doLog(message, fileName, INFO)
}
func (logger *StdLogger) Debug(message string, fileName string) {
	logger.doLog(message, fileName, DEBUG)
}

func (logger *StdLogger) doLog(message string, fileName string, loggingLevel LoggingLevel) {
	if loggingLevels[loggingLevel] <= loggingLevels[logger.loggingLevel] {
		if len(fileName) > 40 {
			fileName = fileName[len(fileName)-40:]
		}
		logger.logger.SetPrefix(time.Now().Format("2006-01-02T15:04:05-07:00") + " " + string(loggingLevel) + " " + strconv.Itoa(logger.pid) + " ")
		logger.logger.Println("--- [" + fmt.Sprintf("%15s", "main") + "] " + fmt.Sprintf("%-40s", fileName) + " : " + message)
		// 2024-11-27T03:19:57.916+05:30  INFO 97914 --- [           main] controller.data.LocationTypeController   : Executed getAllLocationTypes in 1240ms
	}
}