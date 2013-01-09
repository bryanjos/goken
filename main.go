package main

import (
	"fmt"
	"time"
)

const (
	PAGE_SIZE       = 40
	DB_NAME         = "ken"
	JOB_COLLECTION  = "jobs"
	INFO_COLLECTION = "job_info"
	SERVER_NAME     = "localhost"
)

func main() {
	Start()
}

func Start() {
	fmt.Println("Starting")
	for{
		fmt.Println("Getting Jobs")
		jobs, _ := ListJobs()

		fmt.Println("Doing Jobs")
		for i := 0; i < len(jobs); i++ {
			go DoYourJob(jobs[i])
		}
		fmt.Println("Sleeping")
		time.Sleep(5 * time.Minute)
		fmt.Println("Waking")
	}
	fmt.Println("Stopping")
}


func DoYourJob(job Job) {

	info := make(chan []Information)

	tp := TwitterPlugin{}
	fp := FacebookPlugin{}

	go tp.GetData(job, info)
	go fp.GetData(job, info)

	f,p := <-info, <-info

	for i:=0; i < len(f); i++ {
		SaveInformation(f[i])
	}

	for i:=0; i < len(p); i++ {
		SaveInformation(p[i])
	}

	job.Since = time.Now().UTC()

	SaveJob(job)

}
