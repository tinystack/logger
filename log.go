package logger

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
)

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
		w:       []io.Writer{os.Stdout},
		encoder: EncoderJSON,
		caller:  false,
	}
}

func WithLevel(lvl Level) funcOption {
	return func(o *Options) {
		o.lvl = lvl
	}
}

func WithWriter(w io.Writer) funcOption {
	return func(o *Options) {
		o.w = append(o.w, w)
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
	DefaultLogLevel        = DebugLevel
	DefaultLogFile         = "logs/runtime.log"
	DefaultDriver   Driver = newZapDriver
)

func init() {
	rtOutput := &lumberjack.Logger{
		Filename: DefaultLogFile,
	}
	defaultLogger = DefaultDriver(
		WithLevel(DefaultLogLevel),
		WithWriter(os.Stdout),
		WithWriter(rtOutput),
		WithEncoder(EncoderConsole),
		WithCaller(false),
	)
}

func NewLogger(funcOpts ...funcOption) Logger {
	return DefaultDriver(funcOpts...)
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

type NoneLogger struct{}

func (*NoneLogger) Debug(...interface{})          {}
func (*NoneLogger) Info(...interface{})           {}
func (*NoneLogger) Warn(...interface{})           {}
func (*NoneLogger) Error(...interface{})          {}
func (*NoneLogger) Fatal(...interface{})          {}
func (*NoneLogger) Panic(...interface{})          {}
func (*NoneLogger) Debugf(string, ...interface{}) {}
func (*NoneLogger) Infof(string, ...interface{})  {}
func (*NoneLogger) Warnf(string, ...interface{})  {}
func (*NoneLogger) Errorf(string, ...interface{}) {}
func (*NoneLogger) Fatalf(string, ...interface{}) {}
func (*NoneLogger) Panicf(string, ...interface{}) {}
