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

type InformationCollection []Information

func (s InformationCollection) Len() int {
	return len(s)
}
func (s InformationCollection) Less(i, j int) bool {
	return s[i].Time.Unix() < s[j].Time.Unix()
}
func (s InformationCollection) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func ListInformation(slug string, since time.Time) (InformationCollection, error) {
	session, err := mgo.Dial(Configuration.MongoDB_Server)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(Configuration.DB_Name).C(Configuration.Info_Collection)

	results := []Information{}
	err = c.Find(bson.D{{"jobslug", slug}, {"Time", bson.D{{"$lt", since}}}}).Limit(Configuration.Page_Size).All(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func SaveInformation(job Information) error {
	session, err := mgo.Dial(Configuration.MongoDB_Server)
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB(Configuration.DB_Name).C(Configuration.Info_Collection)

	_, err = c.Upsert(bson.M{"id": job.Id}, &job)

	if err != nil {
		return err
	}

	return nil
}
