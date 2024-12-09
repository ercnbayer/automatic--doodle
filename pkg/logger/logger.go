package logger

import (
	"automatic-doodle/pkg/env"
	"automatic-doodle/pkg/utils"
	"automatic-doodle/types"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"time"

	"github.com/fatih/color"
)

type colors func(format string, a ...interface{}) string

type LoggerColors struct {
	Trace   colors
	Info    colors
	Warning colors
	Error   colors
	Fatal   colors
}

var DefaultColors = LoggerColors{
	Trace:   color.CyanString,
	Info:    color.GreenString,
	Warning: color.YellowString,
	Error:   color.RedString,
	Fatal:   color.MagentaString,
}

type LogLevels string

const (
	LOG_LEVEL_TRACE   LogLevels = "TRACE"
	LOG_LEVEL_INFO    LogLevels = "INFO"
	LOG_LEVEL_WARNING LogLevels = "WARNING"
	LOG_LEVEL_ERROR   LogLevels = "ERROR"
	LOG_LEVEL_FATAL   LogLevels = "FATAL"
)

type ProdLog struct {
	App     string    `json:"app"`
	Module  string    `json:"module"`
	Level   LogLevels `json:"level"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

type Logger struct {
	env        types.GoEnv
	dateFormat string
	logFormat  string
	app        string
	module     string
	colors     LoggerColors
}

func (l *Logger) getColor(level LogLevels) colors {
	switch level {
	case LOG_LEVEL_TRACE:
		return l.colors.Trace
	case LOG_LEVEL_INFO:
		return l.colors.Info
	case LOG_LEVEL_WARNING:
		return l.colors.Warning
	case LOG_LEVEL_ERROR:
		return l.colors.Error
	case LOG_LEVEL_FATAL:
		return l.colors.Fatal
	default:
		return color.GreenString
	}
}

func (l *Logger) serializeLog(message string, level LogLevels) string {
	now := time.Now()
	if l.env == types.GoEnvProduction {
		log := ProdLog{
			App:     l.app,
			Module:  l.module,
			Level:   level,
			Message: message,
			Date:    now,
		}
		msg, err := json.Marshal(log)
		if err != nil {
			fmt.Println(err)
		}

		return string(msg)
	}

	color := l.getColor(level)
	msg := color(l.logFormat, now.Format(l.dateFormat), l.app, l.module)

	return fmt.Sprintf(msg+": %s\n", message)
}

func (l *Logger) log(format string, level LogLevels, args ...any) {
	for i, arg := range args {
		if utils.IsStruct(reflect.TypeOf(arg)) || reflect.TypeOf(arg).Kind() == reflect.Map {
			args[i], _ = json.Marshal(arg)
		}
	}

	message := fmt.Sprintf(format, args...)
	log := l.serializeLog(message, level)
	io.WriteString(os.Stdout, log)
}

func (l *Logger) Trace(format string, args ...any) {
	l.log(format, LOG_LEVEL_TRACE, args...)
}

func (l *Logger) Info(format string, args ...any) {
	l.log(format, LOG_LEVEL_INFO, args...)
}

func (l *Logger) Warning(format string, args ...any) {
	l.log(format, LOG_LEVEL_WARNING, args...)
}

func (l *Logger) Error(format string, args ...any) {
	l.log(format, LOG_LEVEL_ERROR, args...)
}

func (l *Logger) Fatal(format string, args ...any) {
	l.log(format, LOG_LEVEL_FATAL, args...)
	os.Exit(1)
}

func New(module string) *Logger {
	return &Logger{
		app:        "AutoDoodle",
		module:     module,
		dateFormat: "2006-01-02T15:04:05.00000000000Z07:00",
		colors:     DefaultColors,
		logFormat:  "[%s] | [%s] - [%s]",
		env:        env.GO_ENV,
	}
}
