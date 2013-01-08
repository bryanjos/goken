package main

import (
	"testing"
	"time"
)

func Test_SaveJob(t *testing.T) {
	err := SaveJob(Job{"Test", "test", time.Now().UTC(), []string{"test"}, [2]float32{30, -90}, 30, time.Now().UTC()})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func Test_ListJobs(t *testing.T) {
	_, err := ListJobs()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func Test_GetJob(t *testing.T) {
	result, err := GetJob("test")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if result.Name != "Test" {
		t.Log("Expected: test, Got: " + result.Name)
		t.Fail()
	}
}
