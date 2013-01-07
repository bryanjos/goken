package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Job struct {
	Name        string
	Slug        string
	Time        time.Time
	Tags        []string
	Coordinates [2]float32
	Distance    float32
	Since       time.Time
}

type Information struct {
	Source      string
	Id          string
	Creator     string
	Time        time.Time
	Data        string
	Coordinates [2]float32
	JobSlug     string
}

func main() {
	start()
}

func start() {
	result, err := AllJobs()
	if err != nil {
		panic(err)
	}

	fmt.Println("Name:", result[0].Name)
}

func AllJobs() ([]Job, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB("ken").C("jobs")

	results := []Job{}
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}

	return results,nil
}

func OneJob(slug string) (Job, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return Job{}, err
	}
	defer session.Close()

	c := session.DB("ken").C("jobs")

	result := Job{}
	err = c.Find(bson.M{"slug": slug}).One(&result)
	if err != nil {
		return Job{}, err
	}

	return result,nil
}
