package factory

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestNewCommonInterface(t *testing.T) {
	t1 := NewCommonInterface(1)
	s1 := t1.Producer("t1")
	assert.Equal(t, s1, "t1"+"生产水果!")
	t2 := NewCommonInterface(2)
	s2 := t2.Producer("t2")
	assert.Equal(t, s2, "t2"+"生产大米!")
}
