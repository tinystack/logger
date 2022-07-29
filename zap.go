package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	zap *zap.SugaredLogger
}

var zapLevel = map[Level]zapcore.Level{
	DebugLevel: zapcore.DebugLevel,
	InfoLevel:  zapcore.InfoLevel,
	WarnLevel:  zapcore.WarnLevel,
	ErrorLevel: zapcore.ErrorLevel,
	PanicLevel: zapcore.PanicLevel,
	FatalLevel: zapcore.FatalLevel,
}

func newZapDriver(funcOpts ...funcOption) Logger {
	opts := defaultOptions()
	for _, f := range funcOpts {
		f(opts)
	}

	lvl := zap.InfoLevel
	if l, ok := zapLevel[opts.lvl]; ok {
		lvl = l
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(lvl)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.FunctionKey = "fn"
	encoderConfig.MessageKey = "msg"

	var encoder zapcore.Encoder
	switch opts.encoder {
	case EncoderConsole:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	var syncer []zapcore.WriteSyncer
	for _, w := range opts.w {
		syncer = append(syncer, zapcore.AddSync(w))
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(syncer...),
		atomicLevel,
	)
	z := zap.New(core, zap.WithCaller(opts.caller), zap.AddCallerSkip(2)).Sugar()
	return &zapLogger{
		zap: z,
	}
}

func (l *zapLogger) z() *zap.SugaredLogger {
	if l.zap == nil {
		panic("logger: Zap.zap not initialized")
	}
	return l.zap
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.z().Debug(args...)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.z().Info(args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.z().Warn(args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.z().Error(args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.z().Fatal(args...)
}

func (l *zapLogger) Panic(args ...interface{}) {
	l.z().Panic(args...)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.z().Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.z().Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.z().Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.z().Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.z().Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.z().Panicf(format, args...)
}

func (l *zapLogger) Debugt(msg string, args T) {
	l.z().Debugw(msg, l.keysAndValues(args)...)
}

func (l *zapLogger) Infot(msg string, args T) {
	l.z().Infow(msg, l.keysAndValues(args)...)
}

func (l *zapLogger) Warnt(msg string, args T) {
	l.z().Warnw(msg, l.keysAndValues(args)...)
}

func (l *zapLogger) Errort(msg string, args T) {
	l.z().Errorw(msg, l.keysAndValues(args)...)
}

func (l *zapLogger) Fatalt(msg string, args T) {
	l.z().Fatalw(msg, l.keysAndValues(args)...)
}

func (l *zapLogger) Panict(msg string, args T) {
	l.z().Panicw(msg, l.keysAndValues(args)...)
}

func (l *zapLogger) keysAndValues(args T) []interface{} {
	var keysAndValues []interface{}
	for k, v := range args {
		keysAndValues = append(keysAndValues, k, v)
	}
	return keysAndValues
}
