package main

import "testing"

func TestGetHello(t *testing.T) {
	if getHello() == "Hello go!" {
		t.Log("Test ok")
	} else {
		t.Log("Test failed")
	}
}
