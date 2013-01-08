package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type tFacebookSearchDummy struct {
	Object tFacebookSearch
}

type tFacebookSearch struct {
	Data []tFacebookSearchResult
}

type tFacebookSearchResult struct {
	Id           string
	From         tFacebookSearchResultFrom
	Created_time string
	Message      string
}

type tFacebookSearchResultFrom struct {
	Id   string
	Name string
}

func Facebook_GetData(job Job) ([]Information, error) {
	tags := strings.Join(job.Tags, " ")
	since := fmt.Sprintf("&since:%d", job.Time.Unix())
	geocode := fmt.Sprintf("&center=%f,%f&distance=%d", job.Coordinates[1], job.Coordinates[0], job.Distance)
	other := "&type=post"
	url := "https://graph.facebook.com/search?q=" + tags + since + geocode + other

	jsonStr := getJsonFromUrl(url)
	var results tFacebookSearchDummy
	err := json.Unmarshal([]uint8(jsonStr), &results)

	if err != nil {
		return []Information{}, err
	}

	info_list := make([]Information, len(results.Object.Data))
	for i := 0; i < len(results.Object.Data); i++ {
		parsed_time, _ := time.Parse("2006-01-02T15:04:05Z", strings.Replace(results.Object.Data[i].Created_time, "+0000", "Z", -1))
		coordinates := [2]float32{0.0, 0.0}

		info := Information{"Facebook", results.Object.Data[i].Id, results.Object.Data[i].From.Name, parsed_time, results.Object.Data[i].Message, coordinates, job.Slug}
		info_list[i] = info
	}

	return info_list, nil
}
