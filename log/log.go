package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	lj "gopkg.in/natefinch/lumberjack.v2"
)

const EVENT_CALLPOLICE = "event:call_police"

type Level = logrus.Level

const (
	// 与logrus.InfoLevel 定义保持一至
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	traceLevel
)

var moduleName = ""

func init() {
	// 默认级别 info
	logrus.SetLevel(InfoLevel)
	logrus.SetFormatter(&MyFormatter{})
}

// Init log service
func Init(debugMode bool, filename string) {
	if debugMode {
		SetLogLevel(DebugLevel)
		SetOutput(filename, true)
	} else {
		SetLogLevel(InfoLevel)
		SetOutput(filename, false)
	}
}

type MyFormatter struct {
	PID int
}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05,000")
	var newLog string = fmt.Sprintf("%s[%s %s]: %s\n", moduleName, LevelString(entry.Level), timestamp, entry.Message)
	b.WriteString(newLog)
	return b.Bytes(), nil
}

func LevelString(level Level) string {
	if b, err := level.MarshalText(); err == nil {
		return strings.ToUpper(string(b))
	} else {
		return "UNKNOWN"
	}
}

var ParseLevl = logrus.ParseLevel

// 设置模块名称
func SetModuleName(module string) {
	moduleName = fmt.Sprintf("%s ", module)
}

// 设置日志级别
func SetLogLevel(level Level) {
	logrus.SetLevel(level)
}

// 设置日志的输出
func SetOutput(filename string, stdout bool) {
	outputs := []io.Writer{}
	if stdout || strings.EqualFold(filename, "stdout") {
		outputs = append(outputs, os.Stdout)
	}

	if len(filename) > 0 && !strings.EqualFold(filename, "stdout") {
		outputs = append(outputs, &lj.Logger{
			Filename:   filename,
			MaxSize:    100,
			MaxBackups: 5,
			MaxAge:     365,
			Compress:   false,
		})
	}

	logrus.SetOutput(io.MultiWriter(outputs...))
}

func Trace(v ...interface{}) {
	logrus.Trace(v...)
}

func Debug(v ...interface{}) {
	logrus.Debug(v...)
}

func Info(v ...interface{}) {
	logrus.Info(v...)
}

func Warning(v ...interface{}) {
	logrus.Warning(v...)
}

func Warn(v ...interface{}) {
	logrus.Warn(v...)
}

func Error(v ...interface{}) {
	logrus.Error(v...)
}

func Panic(v ...interface{}) {
	logrus.Error(append([]interface{}{EVENT_CALLPOLICE}, v...)...)
}

func Fatal(v ...interface{}) {
	logrus.Fatal(v...)
}

func Tracef(format string, v ...interface{}) {
	logrus.Tracef(format, v...)
}

func Debugf(format string, v ...interface{}) {
	logrus.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	logrus.Infof(format, v...)
}

func Warningf(format string, v ...interface{}) {
	logrus.Warningf(format, v...)
}

func Warnf(format string, v ...interface{}) {
	logrus.Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logrus.Errorf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	logrus.Errorf(fmt.Sprintf("%s %s", EVENT_CALLPOLICE, format), v...)
}

func Fatalf(format string, v ...interface{}) {
	logrus.Fatalf(format, v...)
}
