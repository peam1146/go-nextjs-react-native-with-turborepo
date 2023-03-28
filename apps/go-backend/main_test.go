package main

import (
	"pack"
	"testing"
	"time"
)

func TestGet1(t *testing.T) {
	time.Sleep(10 * time.Second)
	if pack.Get1() != 1 {
		t.Errorf("not eq 1")
	}
}
