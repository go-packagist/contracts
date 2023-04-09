package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStdoutLogger(t *testing.T) {
	s := NewStdoutLogger("prod")

	assert.Nil(t, s.Emergency("emergency"))
	assert.Nil(t, s.Alert("alert"))
	assert.Nil(t, s.Critical("critical"))
	assert.Nil(t, s.Error("error"))
	assert.Nil(t, s.Warning("warning"))
	assert.Nil(t, s.Notice("notice"))
	assert.Nil(t, s.Info("info"))
	assert.Nil(t, s.Debug("debug"))

	assert.Nil(t, s.Emergencyf("emergency %s", "test"))
	assert.Nil(t, s.Alertf("alert %s", "test"))
	assert.Nil(t, s.Criticalf("critical %s", "test"))
	assert.Nil(t, s.Errorf("error %s", "test"))
	assert.Nil(t, s.Warningf("warning %s", "test"))
	assert.Nil(t, s.Noticef("notice %s", "test"))
	assert.Nil(t, s.Infof("info %s", "test"))
	assert.Nil(t, s.Debugf("debug %s", "test"))

	assert.Nil(t, s.Log(Emergency, "emergency"))
}
