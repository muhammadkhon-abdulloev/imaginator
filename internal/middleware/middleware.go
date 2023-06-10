package middleware

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/muhammadkhon-abdulloev/pkg/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Middleware struct {
	lg *logger.Logger
}

var (
	Option = fx.Provide(New)
)

type Params struct {
	fx.In
}

func New(lg *logger.Logger) *Middleware {
	return &Middleware{
		lg: lg,
	}
}

func (m *Middleware) UnaryRecoverInterceptor() grpc.UnaryServerInterceptor {
	return recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(_recover))
}

func (m *Middleware) UnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	return logging.UnaryServerInterceptor(_interceptorLogger(m.lg))
}

func (m *Middleware) StreamRecoverInterceptor() grpc.StreamServerInterceptor {
	return recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(_recover))
}

func (m *Middleware) StreamLoggingInterceptor() grpc.StreamServerInterceptor {
	return logging.StreamServerInterceptor(_interceptorLogger(m.lg))
}

func _recover(p any) (err error) {
	return status.Errorf(codes.Internal, "panic error: %v", p)
}

func _interceptorLogger(lg *logger.Logger) logging.LoggerFunc {
	return func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, 2)

		for i := 0; i < len(fields); i += 2 {
			key, ok := fields[i].(string)
			if !ok {
				lg.Error("invalid key type")
				return
			}

			value := fields[i+1]

			switch v := value.(type) {
			case string:
				f = append(f, zap.String(key, v))
			case int:
				f = append(f, zap.Int(key, v))
			case bool:
				f = append(f, zap.Bool(key, v))
			default:
				f = append(f, zap.Any(key, v))
			}
		}

		var level zapcore.Level
		switch lvl {
		case logging.LevelDebug:
			level = zapcore.DebugLevel
		case logging.LevelWarn:
			level = zapcore.WarnLevel
		case logging.LevelError:
			level = zapcore.ErrorLevel
		default:
			level = zapcore.InfoLevel
		}

		lg.Log(level, msg, f...)
	}
}
