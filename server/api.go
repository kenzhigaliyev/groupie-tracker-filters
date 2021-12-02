package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Groupie struct {
	Artists  string `json:"artists"`
	Relation string `json:"relation"`
}

type Artists struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	DatesLocations map[string][]string `json:"datesLocations"`
	Result         bool
	NameCities     map[string]bool
}

type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

var ArtistsNew []Artists

func Func() {
	var Url = "https://groupietrackers.herokuapp.com/api"
	var GroupieNew = Groupie{}
	if !Data(Url, &GroupieNew) {
		ArtistsNew[0].Result = false
		return
	}

	if !Data(GroupieNew.Artists, &ArtistsNew) {
		ArtistsNew[0].Result = false
		return
	}

	var RelationNew = Relation{}
	if !Data(GroupieNew.Relation, &RelationNew) {
		ArtistsNew[0].Result = false
		return
	}
	for index := range ArtistsNew {
		ArtistsNew[index].DatesLocations = RelationNew.Index[index].DatesLocations
	}
	ArtistsNew[0].Result = true
}

func Data(url string, val interface{}) bool {
	res, err := http.Get(url)
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false
	}
	err = json.Unmarshal(body, &val)
	if err != nil {
		return false
	}
	return true
}
