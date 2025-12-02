package main

import "testing"

func TestRepeatedRun(t *testing.T) {
	result := repeatedRun("123123", 3)
	if result != true {
		t.Error("123123, 3 failed")
	}

	result = repeatedRun("123123", 2)
	if result != false {
		t.Error("123123, 2 failed")
	}
}
