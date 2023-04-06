package grpc

import (
	"context"
	"fmt"
	"github.com/sparrow-community/plugins/v4/logger/grpc/proto"
	"net/url"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go-micro.dev/v4/logger"
)

type zapLog struct {
	cfg  zap.Config
	zap  *zap.Logger
	opts logger.Options
	sync.RWMutex
	fields     map[string]interface{}
	grpcWriter *ZapGrpcWriter
}

func (l *zapLog) Init(opts ...logger.Option) error {
	var err error

	for _, o := range opts {
		o(&l.opts)
	}

	if serviceName, ok := l.opts.Context.Value(ServiceNameKey{}).(string); ok {
		if client, ok := l.opts.Context.Value(ClientKey{}).(proto.LoggerService); ok {
			l.grpcWriter = &ZapGrpcWriter{
				serviceName: serviceName,
				client:      client,
				closed:      0,
				closeMutex:  sync.Mutex{},
			}
		}
	}

	zapConfig := zap.NewProductionConfig()
	if config, ok := l.opts.Context.Value(configKey{}).(zap.Config); ok {
		zapConfig = config
	}

	if config, ok := l.opts.Context.Value(encoderConfigKey{}).(zapcore.EncoderConfig); ok {
		zapConfig.EncoderConfig = config
	}
	err = zap.RegisterSink("grpc", func(url *url.URL) (zap.Sink, error) {
		return l.grpcWriter, nil
	})
	zapConfig.OutputPaths = append(zapConfig.OutputPaths, "grpc:")
	// Set log Level if not default
	zapConfig.Level = zap.NewAtomicLevel()
	if l.opts.Level != logger.InfoLevel {
		zapConfig.Level.SetLevel(loggerToZapLevel(l.opts.Level))
	}

	log, err := zapConfig.Build(zap.AddCallerSkip(l.opts.CallerSkipCount))
	if err != nil {
		return err
	}

	// Adding seed fields if exist
	if l.opts.Fields != nil {
		var data []zap.Field
		for k, v := range l.opts.Fields {
			data = append(data, zap.Any(k, v))
		}
		log = log.With(data...)
	}

	// Adding namespace
	if namespace, ok := l.opts.Context.Value(namespaceKey{}).(string); ok {
		log = log.With(zap.Namespace(namespace))
	}

	// Adding options
	if options, ok := l.opts.Context.Value(optionsKey{}).([]zap.Option); ok {
		log = log.WithOptions(options...)
	}

	// defer log.Sync() ??

	l.cfg = zapConfig
	l.zap = log
	l.fields = make(map[string]interface{})

	return nil
}

func (l *zapLog) Fields(fields map[string]interface{}) logger.Logger {
	l.Lock()
	nfields := make(map[string]interface{}, len(l.fields))
	for k, v := range l.fields {
		nfields[k] = v
	}
	l.Unlock()
	for k, v := range fields {
		nfields[k] = v
	}

	data := make([]zap.Field, 0, len(nfields))
	for k, v := range fields {
		data = append(data, zap.Any(k, v))
	}

	zl := &zapLog{
		cfg:    l.cfg,
		zap:    l.zap.With(data...),
		opts:   l.opts,
		fields: make(map[string]interface{}),
	}

	return zl
}

func (l *zapLog) Error(err error) logger.Logger {
	return l.Fields(map[string]interface{}{"error": err})
}

func (l *zapLog) Log(level logger.Level, args ...interface{}) {
	l.RLock()
	data := make([]zap.Field, 0, len(l.fields))
	for k, v := range l.fields {
		data = append(data, zap.Any(k, v))
	}
	l.RUnlock()

	lvl := loggerToZapLevel(level)
	msg := fmt.Sprint(args...)
	switch lvl {
	case zap.DebugLevel:
		l.zap.Debug(msg, data...)
	case zap.InfoLevel:
		l.zap.Info(msg, data...)
	case zap.WarnLevel:
		l.zap.Warn(msg, data...)
	case zap.ErrorLevel:
		l.zap.Error(msg, data...)
	case zap.FatalLevel:
		l.zap.Fatal(msg, data...)
	}
}

func (l *zapLog) Logf(level logger.Level, format string, args ...interface{}) {
	l.RLock()
	data := make([]zap.Field, 0, len(l.fields))
	for k, v := range l.fields {
		data = append(data, zap.Any(k, v))
	}
	l.RUnlock()

	lvl := loggerToZapLevel(level)
	msg := fmt.Sprintf(format, args...)
	switch lvl {
	case zap.DebugLevel:
		l.zap.Debug(msg, data...)
	case zap.InfoLevel:
		l.zap.Info(msg, data...)
	case zap.WarnLevel:
		l.zap.Warn(msg, data...)
	case zap.ErrorLevel:
		l.zap.Error(msg, data...)
	case zap.FatalLevel:
		l.zap.Fatal(msg, data...)
	}
}

func (l *zapLog) String() string {
	return "zap"
}

func (l *zapLog) Options() logger.Options {
	return l.opts
}

// NewLogger New builds a new logger based on options.
func NewLogger(opts ...logger.Option) (logger.Logger, error) {
	// Default options
	options := logger.Options{
		Level:           logger.InfoLevel,
		Fields:          make(map[string]interface{}),
		Out:             os.Stderr,
		Context:         context.Background(),
		CallerSkipCount: 2,
	}

	l := &zapLog{opts: options}
	if err := l.Init(opts...); err != nil {
		return nil, err
	}

	return l, nil
}

func loggerToZapLevel(level logger.Level) zapcore.Level {
	switch level {
	case logger.TraceLevel, logger.DebugLevel:
		return zap.DebugLevel
	case logger.InfoLevel:
		return zap.InfoLevel
	case logger.WarnLevel:
		return zap.WarnLevel
	case logger.ErrorLevel:
		return zap.ErrorLevel
	case logger.FatalLevel:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func zapToLoggerLevel(level zapcore.Level) logger.Level {
	switch level {
	case zap.DebugLevel:
		return logger.DebugLevel
	case zap.InfoLevel:
		return logger.InfoLevel
	case zap.WarnLevel:
		return logger.WarnLevel
	case zap.ErrorLevel:
		return logger.ErrorLevel
	case zap.FatalLevel:
		return logger.FatalLevel
	default:
		return logger.InfoLevel
	}
}
