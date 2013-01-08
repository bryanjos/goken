package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Information struct {
	Source      string
	Id          string
	Creator     string
	Time        time.Time
	Data        string
	Coordinates [2]float32
	JobSlug     string
}

func ListInformation(slug string, since time.Time) ([]Information, error) {
	session, err := mgo.Dial(SERVER_NAME)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(DB_NAME).C(INFO_COLLECTION)

	results := []Information{}
	err = c.Find(bson.D{{"jobslug", slug}, {"Time", bson.D{{"$lt", since}}}}).Limit(PAGE_SIZE).All(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func SaveInformation(job Information) error {
	session, err := mgo.Dial(SERVER_NAME)
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("ken").C(INFO_COLLECTION)

	_, err = c.Upsert(bson.M{"id": job.Id}, &job)

	if err != nil {
		return err
	}

	return nil
}