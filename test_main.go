package main

import "testing"

func Test_AllJobs(t *testing.T) {
	_, err := AllJobs()
	if err == nil {
		t.Log("err")
		t.Fail()
	}
}

func Test_OneJob(t *testing.T) {
	result, err := OneJob("sandwiches")
	if err == nil {
		t.Log("err")
		t.Fail()
	}

	if result.Name != "Sandwiches" {
		t.Log("Expected: Sandwiches, Got: " + result.Name)
		t.Fail()
	}
}
