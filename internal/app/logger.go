package app

import "log"

type Logger interface {
	Printf(format string, args ...any)
	Run() error
	Stop()
}

type StdLogger struct {
	processName string
}

func NewStdLogger(processName string) *StdLogger {
	return &StdLogger{processName: processName}
}

func (l *StdLogger) Printf(format string, args ...any) {
	log.Printf("process=%s "+format, append([]any{l.processName}, args...)...)
}

func (l *StdLogger) Run() error {
	l.Printf("logger ready")
	return nil
}

func (l *StdLogger) Stop() {
	l.Printf("logger stopped")
}
