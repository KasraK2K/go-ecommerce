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
	if strings.Contains(config.AppConfig.StdoutLogs, "debug") {
		l.initLogger("\u001b[34mDebug\u001b[0m:     ", v...)
	}
}

func (l *logger) Info(v ...interface{}) {
	if strings.Contains(config.AppConfig.StdoutLogs, "info") {
		l.initLogger("\u001b[32mINFO\u001b[0m:      ", v...)
	}
}

func (l *logger) Notice(v ...interface{}) {
	if strings.Contains(config.AppConfig.StdoutLogs, "notice") {
		l.initLogger("\u001b[36mNOTICE\u001b[0m:    ", v...)
	}
}

func (l *logger) Warning(v ...interface{}) {
	if strings.Contains(config.AppConfig.StdoutLogs, "warning") {
		l.initLogger("\u001b[33mWARNING\u001b[0m:   ", v...)
	}
}

func (l *logger) Error(v ...interface{}) {
	if strings.Contains(config.AppConfig.StdoutLogs, "error") {
		l.initLogger("\u001b[31mERROR\u001b[0m:     ", v...)
	}
}

func (l *logger) Critical(v ...interface{}) {
	if strings.Contains(config.AppConfig.StdoutLogs, "critical") {
		l.initLogger("\u001b[35mCRITICAL\u001b[0m:  ", v...)
	}
}

func (l *logger) Alert(v ...interface{}) {
	if strings.Contains(config.AppConfig.StdoutLogs, "alert") {
		l.initLogger("\u001b[30;43mALERT\u001b[0m:     ", v...)
	}
}

func (l *logger) Emergency(v ...interface{}) {
	if strings.Contains(config.AppConfig.StdoutLogs, "emergency") {
		l.initLogger("\u001b[37;41mEMERGENCY\u001b[0m: ", v...)
	}
}

/* -------------------------------------------------------------------------- */
/*                                File Loggers                                */
/* -------------------------------------------------------------------------- */
func (l *logger) DebugFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FileLogs, "debug") {
		l.initFileLogger("debug", v...)
	}
}

func (l *logger) InfoFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FileLogs, "info") {
		l.initFileLogger("info", v...)
	}
}

func (l *logger) NoticeFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FileLogs, "notice") {
		l.initFileLogger("notice", v...)
	}
}

func (l *logger) WarningFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FileLogs, "warning") {
		l.initFileLogger("warning", v...)
	}
}

func (l *logger) ErrorFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FileLogs, "error") {
		l.initFileLogger("error", v...)
	}
}

func (l *logger) CriticalFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FileLogs, "critical") {
		l.initFileLogger("critical", v...)
	}
}

func (l *logger) AlertFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FileLogs, "alert") {
		l.initFileLogger("alert", v...)
	}
}

func (l *logger) EmergencyFile(v ...interface{}) {
	if strings.Contains(config.AppConfig.FileLogs, "emergency") {
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
