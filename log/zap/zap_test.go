package zap

import (
	"github.com/jfonseca85/aws-sdk-hello-world/log"
	"testing"

	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	logger := NewLogger(zap.NewExample())

	defer func() { _ = logger.Sync() }()

	zlog := log.NewHelper(logger)

	zlog.Debugw("log", "debug")
	zlog.Infow("log", "info")
	zlog.Warnw("log", "warn")
	zlog.Errorw("log", "error")
}
