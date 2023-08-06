package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"strings"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

type writeHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writeHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writeHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func GetLogger() Logger {
	return Logger{e}
}

func (l Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

type CustomFormatter struct {
	TimestampFormat string
}

func getColorByLevel(level logrus.Level) string {
	switch level {
	case logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel:
		return "\x1b[31m" // Red
	case logrus.WarnLevel:
		return "\x1b[33m" // Yellow
	case logrus.InfoLevel:
		return "\x1b[32m" // Green
	case logrus.DebugLevel, logrus.TraceLevel:
		return "\x1b[36m" // Cyan
	default:
		return "" // No color
	}
}

func getColorReset() string {
	return "\x1b[0m" // Reset color
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.UTC().Format(f.TimestampFormat)
	level := strings.ToUpper(entry.Level.String())
	message := entry.Message
	color := getColorByLevel(entry.Level)
	function := ""
	file := ""
	if entry.Caller != nil {
		function = fmt.Sprintf("%s()", entry.Caller.Function)
		file = fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	}
	logLine := fmt.Sprintf("%s[%s]%s: %s - %s | file=%s func=%s%s\n",
		color, level, getColorReset(), timestamp, message, file, function, getColorReset())
	//logLine := fmt.Sprintf("[%s]: %s - %s | file=%s func=%s\n", level, timestamp, message, file, function)

	return []byte(logLine), nil
}

func init() {
	log := logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &CustomFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err.Error())
	}
	if err != nil {
		panic(err.Error())
	}
	f, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err.Error())
	}
	log.SetOutput(io.Discard)
	log.AddHook(&writeHook{
		Writer:    []io.Writer{f, os.Stdout},
		LogLevels: logrus.AllLevels,
	})
	log.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(log)
}
