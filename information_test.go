package main

import (
	"testing"
	"time"
)

func Test_SaveInformation(t *testing.T) {
	err := SaveInformation(Information{"Test", "test", "information_test", time.Now().UTC(), "test", [2]float32{30, -90}, "test"})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func Test_ListInformation(t *testing.T) {
	_, err := ListInformation("test", time.Now().UTC().Add(-3000))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
