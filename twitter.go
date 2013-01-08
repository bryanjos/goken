package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TwitterPlugin struct{
}

type tTwitterSearchDummy struct {
	Object tTwitterSearch
}

type tTwitterSearch struct {
	Results []tTwitterSearchResult
}

type tTwitterSearchResult struct {
	Id_str     string
	Created_at string
	From_user  string
	Text       string
	Geo        []uint8
}

type tTwitterSearchResultGeo struct {
	Coordinates []float32
}

func (t *TwitterPlugin) GetData(job Job) ([]Information, error) {
	tags := strings.Join(job.Tags, " OR ") + "%20"
	since := fmt.Sprintf("since:%d-%02d-%02d", job.Time.Year(), job.Time.Month(), job.Time.Day())
	geocode := fmt.Sprintf("&geocode=%f,%f,%fmi", job.Coordinates[1], job.Coordinates[0], job.Distance)
	other := "&count=100&lang=en&result_type=recent"
	url := "http://search.twitter.com/search.json?q=" + tags + since + geocode + other

	jsonStr := getJsonFromUrl(url)
	var results tTwitterSearchDummy
	err := json.Unmarshal([]uint8(jsonStr), &results)

	if err != nil {
		return []Information{}, err
	}

	info_list := make([]Information, len(results.Object.Results))
	for i := 0; i < len(results.Object.Results); i++ {
		parsed_time, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 +0000 MST", results.Object.Results[i].Created_at+" GMT")

		coordinates := [2]float32{0.0, 0.0}

		if results.Object.Results[i].Geo != nil {
			var geo tTwitterSearchResultGeo
			err := json.Unmarshal(results.Object.Results[i].Geo, &geo)

			if err == nil {
				coordinates[0] = geo.Coordinates[1]
				coordinates[1] = geo.Coordinates[0]
			}
		}

		info := Information{"Twitter", results.Object.Results[i].Id_str, results.Object.Results[i].From_user, parsed_time, results.Object.Results[i].Text, coordinates, job.Slug}
		info_list[i] = info
	}

	return info_list, nil
}
