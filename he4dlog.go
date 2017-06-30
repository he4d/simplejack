package he4dlog

import (
	"io"
	"log"
)

//Logger contains all Loggers for the various log levels
type Logger struct {
	Trace   *log.Logger
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
}

//NewSingleWriter creates a new Logger with one single output and returns it
func NewSingleWriter(writer io.Writer) *Logger {
	return NewMultiWriter(writer, writer, writer, writer, writer, writer)
}

//NewMultiWriter creates a new Logger with multiple outputs and returns it
func NewMultiWriter(
	traceHandle io.Writer,
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	fatalHandle io.Writer) *Logger {

	trace := log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	debug := log.New(debugHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	info := log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warning := log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	errlog := log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	fatal := log.New(fatalHandle,
		"FATAL: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger := &Logger{
		Trace:   trace,
		Debug:   debug,
		Info:    info,
		Warning: warning,
		Error:   errlog,
		Fatal:   fatal,
	}
	return logger
}
