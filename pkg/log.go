package pkg

import (
	"fmt"
	"log"
	"os"
	"strings"

	"app/config"
)

type logger struct{}

var Logger logger

/* -------------------------------------------------------------------------- */
/*                               Stdout Loggers                               */
/* -------------------------------------------------------------------------- */
func (l *logger) Debug(v ...interface{}) {
	if strings.Contains(config.AppConfig.STDOUT_LOGS, "debug") {
		l.initLogger("\u001b[34mDebug\u001b[0m:     ", v...)
	}
}

func (l *logger) Info(v ...interface{}) {
	if strings.Contains(config.AppConfig.STDOUT_LOGS, "info") {
		l.initLogger("\u001b[32mINFO\u001b[0m:      ", v...)
	}
}

func (l *logger) Notice(v ...interface{}) {
	if strings.Contains(config.AppConfig.STDOUT_LOGS, "notice") {
		l.initLogger("\u001b[36mNOTICE\u001b[0m:    ", v...)
	}
}

func (l *logger) Warning(v ...interface{}) {
	if strings.Contains(config.AppConfig.STDOUT_LOGS, "warning") {
		l.initLogger("\u001b[33mWARNING\u001b[0m:   ", v...)
	}
}

func (l *logger) Error(v ...interface{}) {
	if strings.Contains(config.AppConfig.STDOUT_LOGS, "error") {
		l.initLogger("\u001b[31mERROR\u001b[0m:     ", v...)
	}
}

func (l *logger) Critical(v ...interface{}) {
	if strings.Contains(config.AppConfig.STDOUT_LOGS, "critical") {
		l.initLogger("\u001b[35mCRITICAL\u001b[0m:  ", v...)
	}
}

func (l *logger) Alert(v ...interface{}) {
	if strings.Contains(config.AppConfig.STDOUT_LOGS, "alert") {
		l.initLogger("\u001b[30;43mALERT\u001b[0m:     ", v...)
	}
}

func (l *logger) Emergency(v ...interface{}) {
	if strings.Contains(config.AppConfig.STDOUT_LOGS, "emergency") {
		l.initLogger("\u001b[37;41mEMERGENCY\u001b[0m: ", v...)
	}
}

/* -------------------------------------------------------------------------- */
/*                                File Loggers                                */
/* -------------------------------------------------------------------------- */
func (l *logger) DebugFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FILE_LOGS, "debug") {
		l.initFileLogger("debug", v...)
	}
}

func (l *logger) InfoFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FILE_LOGS, "info") {
		l.initFileLogger("info", v...)
	}
}

func (l *logger) NoticeFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FILE_LOGS, "notice") {
		l.initFileLogger("notice", v...)
	}
}

func (l *logger) WarningFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FILE_LOGS, "warning") {
		l.initFileLogger("warning", v...)
	}
}

func (l *logger) ErrorFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FILE_LOGS, "error") {
		l.initFileLogger("error", v...)
	}
}

func (l *logger) CriticalFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FILE_LOGS, "critical") {
		l.initFileLogger("critical", v...)
	}
}

func (l *logger) AlertFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FILE_LOGS, "alert") {
		l.initFileLogger("alert", v...)
	}
}

func (l *logger) EmergencyFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FILE_LOGS, "emergency") {
		l.initFileLogger("emergency", v...)
	}
}

/* -------------------------------------------------------------------------- */
/*                              Private Functions                             */
/* -------------------------------------------------------------------------- */
func (l *logger) folderExistence() {
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", os.ModePerm)
	}
}

func (l *logger) initLogger(prefix string, v ...interface{}) {
	flags := log.Ltime
	stdoutLogger := log.New(os.Stdout, prefix, flags)
	stdoutLogger.Println(v...)
}

func (l *logger) initFileLogger(fileName string, v ...interface{}) {
	l.folderExistence()
	flags := log.Ldate | log.Ltime
	file, _ := os.OpenFile(fmt.Sprintf("log/%s.log", fileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	fileLogger := log.New(file, "", flags)
	fileLogger.SetOutput(file)
	fileLogger.Println(v...)
}
