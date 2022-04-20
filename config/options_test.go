package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	c := New(
		WithPath("..", "../..", "../../.."),
		WithName("env"),
		WithType("yml"),
	)

	cfg, err := c.Load()
	if err != nil {
		fmt.Errorf(err.Error())
		t.Errorf("Expected: %s, got: %s", "sa-east-1", err.Error())

	}
	cfg.log.Info("Opa sou um log.Info")
	assert.Equal(t, cfg.v.GetString("aws.default_region"), "sa-east-1")

}
