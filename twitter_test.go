package main

import (
	"testing"
	"time"
)

func Test_Twitter_GetData(t *testing.T) {
	info := make(chan InformationCollection)
	job := Job{"Test", "test", time.Now().UTC(), []string{"test"}, [2]float32{-90, 30}, 30, time.Now().UTC()}
	tp := TwitterPlugin{}
	err := tp.GetData(job, info)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
