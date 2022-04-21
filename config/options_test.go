package config

import (
	"fmt"
	zlog "github.com/jfonseca85/aws-sdk-hello-world/log/zap"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

type testJSONSource struct {
	data string
	sig  chan struct{}
	err  chan struct{}
}

func newTestJSONSource(data string) *testJSONSource {
	return &testJSONSource{data: data, sig: make(chan struct{}), err: make(chan struct{})}
}

func TestConfig(t *testing.T) {

	logger := zlog.NewLogger(zap.NewExample())

	defer func() { _ = logger.Sync() }()

	c := New(
		WithPath("..", "../..", "../../.."),
		WithName("env"),
		WithType("yml"),
		WithLogger(logger),
	)

	cfg, err := c.Load()
	if err != nil {
		fmt.Errorf(err.Error())
		t.Errorf("Expected: %s, got: %s", "sa-east-1", err.Error())

	}
	cfg.log.Info("Opa sou um log.Info")
	cfg.log.Debug("Opa sou um log.Debug")
	assert.Equal(t, cfg.v.GetString("aws.default_region"), "sa-east-1")

}
