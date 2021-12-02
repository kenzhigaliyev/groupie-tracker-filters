package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Error struct {
	Str    string
	number int
}

// index function processes page index
func index(w http.ResponseWriter, r *http.Request) {
	if !ArtistsNew[0].Result {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}

	if r.Method != "GET" {
		Err("405 Method Not Allowed", http.StatusMethodNotAllowed, w, r)
		return
	}

	if r.URL.Path != "/" {
		Err("404 Status Not Found", http.StatusNotFound, w, r)
		return
	}

	t, err := template.ParseFiles("static/templates/index.html")
	if err != nil {
		fmt.Println(err)
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}

	SortNameCities()

	err = t.ExecuteTemplate(w, "index.html", ArtistsNew)
	if err != nil {
		log.Println("Error when parsing a template:", err)
		fmt.Fprintf(w, err.Error())
		return
	}
}

// artist function processes page artist
func artist(w http.ResponseWriter, r *http.Request) {

	if !ArtistsNew[0].Result {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}

	if len(r.URL.Path) < 10 || r.URL.Path[:9] != "/artists/" {
		Err("400 Bad Request", http.StatusBadRequest, w, r)
		return
	}

	if r.Method != "GET" {
		Err("405 Method Not Allowed", http.StatusMethodNotAllowed, w, r)
		return
	}

	val, err := template.ParseFiles("static/templates/artist.html")
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}

	name := strings.TrimPrefix(r.URL.Path, "/artists/")
	id, err1 := strconv.Atoi(name)
	if err1 != nil {
		Err("400 Bad Request", http.StatusBadRequest, w, r)
		return
	}

	if id < 1 {
		Err("400 Bad Request", http.StatusBadRequest, w, r)
		return
	}

	if id > len(ArtistsNew) {
		Err("404 Not Found", http.StatusNotFound, w, r)
		return
	}

	err = val.ExecuteTemplate(w, "artist.html", ArtistsNew[id-1])
	if err != nil {
		log.Println("Error when parsing a template: %s", err)
		fmt.Fprintf(w, err.Error())
		return
	}
}

// filter function processes page filter
func Filter(w http.ResponseWriter, r *http.Request) {

	ArtistsFilter := []Artists{}

	if !ArtistsNew[0].Result {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}

	if r.URL.Path != "/filters/" {
		Err("400 Bad Request", http.StatusBadRequest, w, r)
		return
	}

	if r.Method != "GET" {
		Err("405 Method Not Allowed", http.StatusMethodNotAllowed, w, r)
		return
	}

	CreationDate := r.FormValue("CreationDate")
	if CreationDate == "on" {
		CreationDateFrom := r.FormValue("CreationDateFrom")
		CreationDateTo := r.FormValue("CreationDateTo")

		if CreationDateFrom == "" {
			CreationDateFrom = "0"
		}

		if CreationDateTo == "" {
			CreationDateTo = "2111"
		}

		CDF, err1 := strconv.Atoi(CreationDateFrom)
		CDT, err2 := strconv.Atoi(CreationDateTo)
		if err1 != nil || err2 != nil {
			Err("400 Bad Request", http.StatusBadRequest, w, r)
			return
		}

		if !CheckValue(CDF, CDT) {
			Err("400 Bad Request", http.StatusBadRequest, w, r)
			return
		}

		ArtistsFilter = CheckOnCreationDate(ArtistsFilter, CDF, CDT)
		if len(ArtistsFilter) == 0 {
			Err("Not Found", http.StatusOK, w, r)
			return
		}
	}

	FirstAlbumDate := r.FormValue("FirstAlbumDate")
	var err1 bool
	if FirstAlbumDate == "on" {
		FirstAlbumDateFrom := r.FormValue("FirstFrom")
		FirstAlbumDateTo := r.FormValue("FirstTo")

		if FirstAlbumDateFrom == "" {
			FirstAlbumDateFrom = "20-01-1000"
		}

		if FirstAlbumDateTo == "" {
			FirstAlbumDateTo = "20-01-2111"
		}

		if !CheckValueDate(FirstAlbumDateFrom, FirstAlbumDateTo) {
			Err("400 Bad Request", http.StatusBadRequest, w, r)
			return
		}

		ArtistsFilter, err1 = CheckFirstAlbumDate(ArtistsFilter, FirstAlbumDateFrom, FirstAlbumDateTo)
		if !err1 {
			Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
			return
		}

		if len(ArtistsFilter) == 0 {
			Err("Not Found", http.StatusOK, w, r)
			return
		}
	}

	NumberOfMembers := r.FormValue("NOM")
	if NumberOfMembers == "on" {

		NumberOfMembersFrom := r.FormValue("NOMfrom")
		NumberOfMembersTo := r.FormValue("NOMto")
		if NumberOfMembersFrom == "" {
			NumberOfMembersFrom = "1"
		}

		if NumberOfMembersTo == "" {
			NumberOfMembersTo = "111"
		}

		NOMF, err1 := strconv.Atoi(NumberOfMembersFrom)
		NOMT, err2 := strconv.Atoi(NumberOfMembersTo)
		if err1 != nil || err2 != nil {
			Err("400 Bad Request", http.StatusBadRequest, w, r)
			return
		}

		if !CheckValue(NOMF, NOMT) {
			Err("400 Bad Request", http.StatusBadRequest, w, r)
			return
		}

		ArtistsFilter = CheckOnNumberOfMembers(ArtistsFilter, NOMF, NOMT)
		if len(ArtistsFilter) == 0 {
			Err("Not Found", http.StatusOK, w, r)
			return
		}
	}

	LocationOfConcerts := r.FormValue("LocationOfConcerts")
	if LocationOfConcerts == "on" {
		LocationOfConcertsValue := r.FormValue("LOC")

		if LocationOfConcertsValue != "" {
			ArtistsFilter = CheckOnLocationOfConcerts(ArtistsFilter, LocationOfConcertsValue)
			if len(ArtistsFilter) == 0 {
				Err("Not Found", http.StatusOK, w, r)
				return
			}
		}
	}

	val, err := template.ParseFiles("static/templates/filter.html")
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w, r)
		return
	}

	if len(ArtistsFilter) == 0 {
		err = val.ExecuteTemplate(w, "filter.html", ArtistsNew)
		if err != nil {
			log.Println("Error when parsing a template: %s", err)
			fmt.Fprintf(w, err.Error())
			return
		}
		return
	}

	err = val.ExecuteTemplate(w, "filter.html", ArtistsFilter)
	if err != nil {
		log.Println("Error when parsing a template: %s", err)
		fmt.Fprintf(w, err.Error())
		return
	}
}

// The Err function displays the result of errors
func Err(Str string, Status int, w http.ResponseWriter, r *http.Request) {

	Info := Error{Str, Status}
	val, err := template.ParseFiles("static/templates/error.html")

	if err != nil {
		log.Println("Error when parsing a template: %s", err)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(Status)
	err = val.ExecuteTemplate(w, "error.html", Info)
	if err != nil {
		log.Println("Error when parsing a template: %s", err)
		fmt.Fprintf(w, err.Error())
		return
	}
}

func HandleFuncOwn() {
	http.HandleFunc("/", index)
	http.HandleFunc("/artists/", artist)
	http.HandleFunc("/filters/", Filter)
	log.Println(http.ListenAndServe(":8080", nil))
}
