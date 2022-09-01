package logger

import (
	"io"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
)

type T map[string]interface{}

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debugt(msg string, args T)
	Infot(msg string, args T)
	Warnt(msg string, args T)
	Errort(msg string, args T)
	Fatalt(msg string, args T)
	Panict(msg string, args T)
}

type Level int8

type Driver func(...funcOption) Logger

const (
	NoneLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

var unmarshalLevelText = map[string]Level{
	"debug": DebugLevel,
	"info":  InfoLevel,
	"warn":  WarnLevel,
	"error": ErrorLevel,
	"panic": PanicLevel,
	"fatal": FatalLevel,
}

func ParseLevel(text string) Level {
	text = strings.ToLower(text)
	lvl, _ := unmarshalLevelText[text]
	return lvl
}

type Options struct {
	lvl     Level
	w       []io.Writer
	encoder string
	caller  bool
}

type funcOption func(*Options)

func defaultOptions() *Options {
	return &Options{
		lvl:     DebugLevel,
		encoder: EncoderJSON,
		caller:  false,
	}
}

func WithLevel(lvl Level) funcOption {
	return func(o *Options) {
		o.lvl = lvl
	}
}

func WithWriter(w ...io.Writer) funcOption {
	return func(o *Options) {
		o.w = w
	}
}

func WithCaller(caller bool) funcOption {
	return func(o *Options) {
		o.caller = caller
	}
}

const (
	EncoderJSON    = "json"
	EncoderConsole = "console"
)

func WithEncoder(encoder string) funcOption {
	return func(o *Options) {
		o.encoder = encoder
	}
}

var defaultLogger Logger

var (
	defaultLogLevel        = DebugLevel
	defaultLogFile         = "logs/runtime.log"
	defaultDriver   Driver = newZapDriver
)

func init() {
	rtOutput := &lumberjack.Logger{
		Filename: defaultLogFile,
	}
	defaultLogger = defaultDriver(
		WithLevel(defaultLogLevel),
		WithWriter(rtOutput, os.Stdout),
		WithEncoder(EncoderConsole),
		WithCaller(false),
	)
}

func NewLogger(funcOpts ...funcOption) Logger {
	return defaultDriver(funcOpts...)
}

func DefaultLogger() Logger {
	return defaultLogger
}

func UpdateDefaultLogger(l Logger) {
	defaultLogger = l
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

func Debugt(msg string, args T) {
	defaultLogger.Debugt(msg, args)
}

func Infot(msg string, args T) {
	defaultLogger.Infot(msg, args)
}

func Warnt(msg string, args T) {
	defaultLogger.Warnt(msg, args)
}

func Errort(msg string, args T) {
	defaultLogger.Errort(msg, args)
}

func Fatalt(msg string, args T) {
	defaultLogger.Fatalt(msg, args)
}

func Panict(msg string, args T) {
	defaultLogger.Panict(msg, args)
}

type NoneLogger struct{}

func (*NoneLogger) Debug(args ...interface{})                 {}
func (*NoneLogger) Info(args ...interface{})                  {}
func (*NoneLogger) Warn(args ...interface{})                  {}
func (*NoneLogger) Error(args ...interface{})                 {}
func (*NoneLogger) Fatal(args ...interface{})                 {}
func (*NoneLogger) Panic(args ...interface{})                 {}
func (*NoneLogger) Debugf(format string, args ...interface{}) {}
func (*NoneLogger) Infof(format string, args ...interface{})  {}
func (*NoneLogger) Warnf(format string, args ...interface{})  {}
func (*NoneLogger) Errorf(format string, args ...interface{}) {}
func (*NoneLogger) Fatalf(format string, args ...interface{}) {}
func (*NoneLogger) Panicf(format string, args ...interface{}) {}
func (*NoneLogger) Debugt(msg string, args T)                 {}
func (*NoneLogger) Infot(msg string, args T)                  {}
func (*NoneLogger) Warnt(msg string, args T)                  {}
func (*NoneLogger) Errort(msg string, args T)                 {}
func (*NoneLogger) Fatalt(msg string, args T)                 {}
func (*NoneLogger) Panict(msg string, args T)                 {}
