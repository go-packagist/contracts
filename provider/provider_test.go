package provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testProvider struct {
	*UnimplementedProvider
}

func TestPrivider(t *testing.T) {
	test := &testProvider{}
	_, ok := interface{}(test).(Provider)

	assert.True(t, ok)

	assert.NotPanics(t, func() {
		test.Register()
		test.Boot()
	})
}
