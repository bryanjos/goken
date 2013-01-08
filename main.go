package main

import (
	"fmt"
)

const (
	PAGE_SIZE       = 40
	DB_NAME         = "ken"
	JOB_COLLECTION  = "jobs"
	INFO_COLLECTION = "job_info"
	SERVER_NAME     = "localhost"
)

func main() {
	start()
}

func start() {
	fmt.Println("Starting")

	fmt.Println("Stopping")
}
