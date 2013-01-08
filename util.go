package main

import (
	"io/ioutil"
	"net/http"
)

func getJsonFromUrl(url_ string) string {
	r, error := httpGet(url_)
	if error != nil {
		return ""
	}

	data, err := parseResponse(r)
	data = fixBrokenJson(data)
	if err != nil {
		return ""
	}

	return data
}

func httpGet(url_ string) (*http.Response, error) {
	var r *http.Response
	var err error

	r, err = http.Get(url_)

	return r, err
}

func parseResponse(response *http.Response) (string, error) {
	var b []byte
	b, _ = ioutil.ReadAll(response.Body)
	response.Body.Close()
	bStr := string(b)

	return bStr, nil
}

func fixBrokenJson(j string) string { return `{"object":` + j + "}" }
