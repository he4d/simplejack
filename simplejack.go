package simplejack

import (
	"io"
	"io/ioutil"
	"log"
)

//LogLevel is the type for the LogLevels
type LogLevel uint8

const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

//Logger contains all loggers for the various log levels
type Logger struct {
	Trace   *log.Logger
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
}

//New creates a new Logger with and returns it.
//It is necessary to pass a minimum loglevel and a writer to where
//the log should be written to
func New(minLogLevel LogLevel, writer io.Writer) *Logger {
	traceHandle := ioutil.Discard
	debugHandle := ioutil.Discard
	infoHandle := ioutil.Discard
	warningHandle := ioutil.Discard
	errorHandle := ioutil.Discard
	fatalHandle := ioutil.Discard
	switch minLogLevel {
	case TRACE:
		traceHandle = writer
		debugHandle = writer
		infoHandle = writer
		warningHandle = writer
		errorHandle = writer
		fatalHandle = writer
	case DEBUG:
		debugHandle = writer
		infoHandle = writer
		warningHandle = writer
		errorHandle = writer
		fatalHandle = writer
	case INFO:
		infoHandle = writer
		warningHandle = writer
		errorHandle = writer
		fatalHandle = writer
	case WARNING:
		warningHandle = writer
		errorHandle = writer
		fatalHandle = writer
	case ERROR:
		errorHandle = writer
		fatalHandle = writer
	case FATAL:
		fatalHandle = writer
	}
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
