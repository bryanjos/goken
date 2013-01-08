package main

import (
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

func ListJobs() ([]Job, error) {
	session, err := mgo.Dial(SERVER_NAME)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(DB_NAME).C(JOB_COLLECTION)

	results := []Job{}
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetJob(slug string) (Job, error) {
	session, err := mgo.Dial(SERVER_NAME)
	if err != nil {
		return Job{}, err
	}
	defer session.Close()

	c := session.DB(DB_NAME).C(JOB_COLLECTION)

	result := Job{}
	err = c.Find(bson.M{"slug": slug}).One(&result)
	if err != nil {
		return Job{}, err
	}

	return result, nil
}

func SaveJob(job Job) error {
	session, err := mgo.Dial(SERVER_NAME)
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB(DB_NAME).C(JOB_COLLECTION)

	_, err = c.Upsert(bson.M{"slug": job.Slug}, &job)

	if err != nil {
		return err
	}

	return nil
}
