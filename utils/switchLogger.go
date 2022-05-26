package utils

import "log"

var (
	DefaultSwitchLogger = &SwitchLogger{}
)

type SwitchLogger struct {
	showLog bool
}

func (l *SwitchLogger) SetShowLog(showLog bool) {
	l.showLog = showLog
}

func (l *SwitchLogger) Printf(format string, v ...interface{}) {
	if !l.showLog {
		return
	}
	log.Printf(format, v...)
}

func (l *SwitchLogger) Print(v ...interface{}) {
	if !l.showLog {
		return
	}
	log.Print(v...)
}

func (l *SwitchLogger) Println(v ...interface{}) {
	if !l.showLog {
		return
	}
	log.Println(v...)
}
