package provider

import "testing"

type testProvider struct {
	*UnimplementedProvider
}

func TestPrivider(t *testing.T) {
	if _, ok := interface{}(&testProvider{}).(Provider); !ok {
		t.Error("testProvider does not implement Provider")
	}
}
