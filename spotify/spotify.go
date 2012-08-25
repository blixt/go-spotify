package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Album struct {
	Name         string
	Href         string
	Released     string
	Availability interface{}
}

type Api struct {
}

type Artist struct {
	Name string
	Href string
}

type ExternalId struct {
	Type string
	Id   string
}

type Info struct {
	NumResults int `json:"num_results"`
	Limit      int
	Offset     int
	Query      string
	Type       string
	Page       int
}

type Track struct {
	Name        string
	Href        string
	Artists     []Artist
	Album       Album
	TrackNumber string `json:"track-number"`
	Length      float32
	Popularity  string
	ExternalIds []ExternalId `json:"external-ids"`
}

type query interface {
	GetCallInfo() (path string, params map[string][]string)
	setError(error)
}

type queryBase struct {
	Error error
}

func (query *queryBase) setError(err error) {
	query.Error = err
}

type SearchTrackQuery struct {
	queryBase
	Query  string
	Info   Info
	Tracks []Track
}

func (query *SearchTrackQuery) GetCallInfo() (path string, params map[string][]string) {
	path = "/search/1/track.json"
	params = map[string][]string{
		"q": {query.Query}}
	return
}

const (
	BASE_URL = "http://ws.spotify.com"
)

func GetApi() *Api {
	return &Api{}
}

func (api *Api) call(query query) {
	path, params := query.GetCallInfo()

	v := url.Values{}
	for key, values := range params {
		for _, value := range values {
			v.Add(key, value)
		}
	}

	resp, err := http.Get(fmt.Sprintf("%s%s?%s", BASE_URL, path, v.Encode()))
	if err != nil {
		query.setError(err)
		return
	}
	defer resp.Body.Close()

	if data, err := ioutil.ReadAll(resp.Body); err != nil {
		query.setError(err)
	} else {
		json.Unmarshal(data, &query)
	}
}

func (api *Api) SearchTrack(q string, out chan *SearchTrackQuery) {
	query := &SearchTrackQuery{
		Query: q}

	api.call(query)
	out <- query
}
