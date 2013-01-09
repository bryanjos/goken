package main

import (
	"fmt"
	"time"
	"sort"
	"labix.org/v2/mgo"
)


const (
	PAGE_SIZE       = 40
	DB_NAME         = "ken"
	JOB_COLLECTION  = "jobs"
	INFO_COLLECTION = "job_info"
	MONGODB_SERVER  = "localhost"
	SERVER_PORT 	= ":3000"

)

func main() {
	go StartCollector()
	StartServer()
}

func StartCollector() {
	fmt.Println("Starting")
	CreateIndexes()
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

	info := make(chan InformationCollection)

	twitterPlugin := TwitterPlugin{}
	facebookPlugin := FacebookPlugin{}

	go twitterPlugin.GetData(job, info)
	go facebookPlugin.GetData(job, info)

	results := append(<- info, <- info...)

	sort.Sort(results)

	for i:=0; i < len(results); i++ {
		SaveInformation(results[i])
	}

	job.Since = time.Now().UTC()

	SaveJob(job)

}

func CreateIndexes(){
	session, _ := mgo.Dial(MONGODB_SERVER)
	defer session.Close()

	c := session.DB(DB_NAME).C(JOB_COLLECTION)

	index := mgo.Index{
		Key: []string{"slug"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
	_ = c.EnsureIndex(index)

	c = session.DB(DB_NAME).C(INFO_COLLECTION)

	index = mgo.Index{
		Key: []string{"jobslug", "-time"},
		Unique: false,
		DropDups: false,
		Background: true,
		Sparse: true,
	}
	_ = c.EnsureIndex(index)
}
