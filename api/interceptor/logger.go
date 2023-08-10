package interceptor

import (
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LogModeDev        = "dev"
	LogModeProduction = "prod"
)

// ZapLogger is a middleware and zap to provide an "access log" like logging for each request.
func ZapLogger(logMode string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger, err := NewLogger(logMode)
			if err != nil {
				c.Error(err)
			}
			start := time.Now()

			req := c.Request()
			res := c.Response()
			fields := []zapcore.Field{
				zap.String("remote_ip", c.RealIP()),
				zap.String("latency", time.Since(start).String()),
				zap.String("host", req.Host),
				zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.Int64("size", res.Size),
				zap.String("user_agent", req.UserAgent()),
			}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			logger.Info("Recv request", fields...)

			logger = logger.With(zap.String("request_id", id))

			c.Set("log", logger)
			err = next(c)
			if err != nil {
				c.Error(err)
			}

			n := res.Status
			switch {
			case n >= 500:
				logger.With(zap.Error(err), zap.Int("status", n)).Error("Server error")
			case n >= 400:
				logger.With(zap.Error(err), zap.Int("status", n)).Warn("Client error")
			case n >= 300:
				logger.With(zap.Int("status", n)).Info("Redirection")
			default:
				logger.With(zap.Int("status", n)).Info("Success")
			}

			return nil
		}
	}
}

func NewLogger(serverMode string) (*zap.Logger, error) {

	loggerCfg := &zap.Config{
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	switch serverMode {
	case LogModeProduction:
		loggerCfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		loggerCfg.EncoderConfig = encoderProdConfig
	case LogModeDev:
		loggerCfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		loggerCfg.EncoderConfig = encoderDevConfig
	default:
		loggerCfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		loggerCfg.EncoderConfig = encoderDevConfig
	}

	plain, err := loggerCfg.Build(zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		plain = zap.NewNop()
	}
	return plain, nil
}

var encoderProdConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "severity",
	NameKey:        "logger",
	CallerKey:      "caller",
	FunctionKey:    "function",
	MessageKey:     "message",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    encodeLevel(),
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.MillisDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
	EncodeName:     ShortNameEncoder,
}

var encoderDevConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "severity",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "message",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    encodeLevel(),
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.MillisDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
	//EncodeName:     ShortNameEncoder,
}

func ShortNameEncoder(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
	loggerNameArr := strings.Split(loggerName, "/")
	fmt.Println()
	enc.AppendString(loggerNameArr[len(loggerNameArr)-1])
}

func encodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		}
	}
}
