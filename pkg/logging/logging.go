package logging

import (
	"fmt"
	"io"
	"os"

	"github.com/KyberNetwork/cclog/lib/client"
	"github.com/TheZeroSlave/zapsentry"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logLevelInfo  = "info"
	logLevelWarn  = "warn"
	logLevelError = "error"
	logLevelFatal = "fatal"
)

type syncer interface {
	Sync() error
}

// NewFlusher creates a new syncer from given syncer that log a error message if failed to sync.
func NewFlusher(s syncer) func() {
	return func() {
		// ignore the error as the sync function will always fail in Linux
		// https://github.com/uber-go/zap/issues/370
		_ = s.Sync()
	}
}

// newLogger creates a new logger instance.
// The type of logger instance will be different with different application running modes.
func newLogger(logLevel, cclogAddr, cclogName string) (*zap.Logger, zap.AtomicLevel) {
	writers := []io.Writer{os.Stdout}
	if cclogAddr != "" && cclogName != "" {
		ccw := client.NewAsyncLogClient(cclogName, cclogAddr, func(err error) {
			fmt.Fprintln(os.Stdout, "send log error", err)
		})

		writers = append(writers, &UnescapeWriter{w: ccw})
	}

	w := io.MultiWriter(writers...)
	atom, err := zap.ParseAtomicLevel(logLevel)
	if err != nil {
		atom = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	config.CallerKey = "caller"

	encoder := zapcore.NewJSONEncoder(config)
	cc := zap.New(zapcore.NewCore(encoder, zapcore.AddSync(w), atom), zap.AddCaller())

	return cc, atom
}

// NewLogger creates a new logger and a flush function. The flush function should be
// called by consumer before quitting the application.
func NewLogger(
	logLevel, cclogAddr, cclogName, sentryDSN, sentryLevel string,
) (*zap.Logger, zap.AtomicLevel, func(), error) {
	logger, atom := newLogger(logLevel, cclogAddr, cclogName)

	// init sentry if flag dsn exists
	if len(sentryDSN) != 0 {
		sentryClient, err := sentry.NewClient(
			sentry.ClientOptions{
				Dsn: sentryDSN,
			},
		)
		if err != nil {
			return nil, atom, nil, fmt.Errorf("failed to init sentry client: %w", err)
		}

		cfg := zapsentry.Configuration{
			DisableStacktrace: false,
		}

		switch sentryLevel {
		case logLevelInfo:
			cfg.Level = zapcore.InfoLevel
		case logLevelWarn:
			cfg.Level = zapcore.WarnLevel
		case logLevelError:
			cfg.Level = zapcore.ErrorLevel
		case logLevelFatal:
			cfg.Level = zapcore.FatalLevel
		default:
			return nil, atom, nil, fmt.Errorf("invalid log level %v", sentryLevel)
		}

		core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromClient(sentryClient))
		if err != nil {
			return nil, atom, nil, fmt.Errorf("failed to init zap sentry: %w", err)
		}
		// attach to logger core
		logger = zapsentry.AttachCoreToLogger(core, logger)
	}

	return logger, atom, NewFlusher(logger), nil
}
