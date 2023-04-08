package provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testProvider struct {
	*UnimplementedProvider
}

func TestPrivider(t *testing.T) {
	_, ok := interface{}(&testProvider{}).(Provider)

	assert.True(t, ok)
}
