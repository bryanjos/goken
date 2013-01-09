package main

import (
	"testing"
	"time"
)

func Test_Facebook_GetData(t *testing.T) {
	info := make(chan []Information)
	job := Job{"Test", "test", time.Now().UTC(), []string{"test"}, [2]float32{-90, 30}, 30, time.Now().UTC()}
	fp := FacebookPlugin{}
	err := fp.GetData(job, info)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
