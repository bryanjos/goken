package main

import (
	"testing"
	"time"
)

func Test_Facebook_GetData(t *testing.T) {
	job := Job{"Test", "test", time.Now().UTC(), []string{"test"}, [2]float32{-90, 30}, 30, time.Now().UTC()}
	fp := FacebookPlugin{}
	_, err := fp.GetData(job)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
