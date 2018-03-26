package fetch

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Depado/gomonit/models"
)

// MONITROOT is the root URL of the the gomonit service
var MONITROOT = "https://monit.depado.eu"

// LightService is a lighter representation of a running service
// intented to be used in compliance with gomonit's models.Service
type LightService struct {
	Name         string
	URL          string
	ShortURL     string
	RepoURL      string
	Host         string
	BuildURL     string
	LastBuild    models.Build
	LastCommit   models.Commit
	Status       int
	Icon         string
	RepoStars    int
	RepoForks    int
	RepoWatchers int
	Description  string
}

// Current is the actual running services that are displayed on the website
var Current []LightService

// Do actually fetches the data source and parses it
func Do() error {
	var err error
	var resp *http.Response
	var all models.Services

	if resp, err = http.Get(MONITROOT + "/api/dump/own"); err != nil {
		return err
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&all); err != nil {
		return err
	}

	Current = []LightService{}
	for _, s := range all {
		new := LightService{
			Name:     s.Name,
			URL:      s.URL,
			ShortURL: s.ShortURL,
			Host:     s.Host,
			Status:   s.Status,
			Icon:     s.Icon,
		}
		if s.Repo != nil {
			new.RepoURL = s.Repo.URL
			new.RepoStars = s.Repo.Stars
			new.RepoForks = s.Repo.Forks
			new.RepoWatchers = s.Repo.Watchers
		}
		if s.CI != nil {
			new.BuildURL = s.CI.URL
		}
		if len(s.LastBuilds) > 0 {
			new.LastBuild = s.LastBuilds[0]
		}
		if len(s.LastCommits) > 0 {
			new.LastCommit = s.LastCommits[0]
		}
		Current = append(Current, new)
	}
	return nil
}

// Start starts the fetching of the data and retrieves data once every hour
func Start() {
	var err error

	tc := time.NewTicker(30 * time.Minute)
	for {
		if err = Do(); err != nil {
			log.Println(err)
		}
		<-tc.C
	}
}
