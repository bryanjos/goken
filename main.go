package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"sort"
	"time"
)

func main() {
	Configuration = InitConfig()
	go Collect()
	StartServer()
}

func Collect() {
	fmt.Println("Starting")
	CreateIndexes()
	for {
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

	results := append(<-info, <-info...)

	sort.Sort(results)

	for i := 0; i < len(results); i++ {
		SaveInformation(results[i])
	}

	job.Since = time.Now().UTC()

	SaveJob(job)

}

func CreateIndexes() {
	session, _ := mgo.Dial(Configuration.MongoDB_Server)
	defer session.Close()

	c := session.DB(Configuration.DB_Name).C(Configuration.Job_Collection)

	index := mgo.Index{
		Key:        []string{"slug"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	_ = c.EnsureIndex(index)

	c = session.DB(Configuration.DB_Name).C(Configuration.Job_Collection)

	index = mgo.Index{
		Key:        []string{"jobslug", "-time"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}
	_ = c.EnsureIndex(index)
}
