package singletonpattern

import (
	"testing"
)

func TestInstance(t *testing.T) {
	t2 := Instance()
	t1 := Instance()
	if t1 != t2 {
		t.Errorf("t1 = %p, t2 =  %p", t1, t2)
	}
}

func TestInstanceLazy(t *testing.T) {
	t1 := InstanceLazy()
	t2 := InstanceLazy()
	if t1 != t2 {
		t.Errorf("t1 = %p, t2 =  %p", t1, t2)
	}
}
