package server

import (
	"strconv"
	"strings"
	"time"
)

//SortNameCities function sort name's cities where artists was
func SortNameCities() {

	ArtistsNew[0].NameCities = make(map[string]bool)
	for _, value := range ArtistsNew {
		for key := range value.DatesLocations {
			ArtistsNew[0].NameCities[key] = true
		}
	}
}

// CheckValue function check on valid input date
func CheckValue(From, To int) bool {

	if From > To {
		return false
	}

	if From < 0 || To < 0 {
		return false
	}
	return true
}

// CheckOnCreationDate function check to correspondence Creation Date of Artists
func CheckOnCreationDate(ArtistFilter []Artists, From, To int) []Artists {

	if len(ArtistFilter) > 0 {
		FilterInFilter := []Artists{}
		for _, value := range ArtistFilter {

			if From <= value.CreationDate && value.CreationDate <= To {
				FilterInFilter = append(FilterInFilter, value)
			}
		}
		return FilterInFilter
	}

	for _, value := range ArtistsNew {

		if From <= value.CreationDate && value.CreationDate <= To {
			ArtistFilter = append(ArtistFilter, value)
		}
	}

	return ArtistFilter
}

// CheckOnNumberOfMembers function check to correspondence Creation Date of Artists
func CheckOnNumberOfMembers(ArtistFilter []Artists, From, To int) []Artists {

	if len(ArtistFilter) > 0 {
		FilterInFilter := []Artists{}
		for _, value := range ArtistFilter {

			if From <= len(value.Members) && len(value.Members) <= To {
				FilterInFilter = append(FilterInFilter, value)
			}
		}
		return FilterInFilter
	}

	for _, value := range ArtistsNew {

		if From <= len(value.Members) && len(value.Members) <= To {
			ArtistFilter = append(ArtistFilter, value)
		}
	}

	return ArtistFilter
}

// CheckValueDate function check on valid date
func CheckValueDate(From, To string) bool {

	from := strings.Split(From, "-")
	to := strings.Split(To, "-")

	if len(from) != 3 || len(to) != 3 {
		return false
	}

	if !checkDate(from) {
		return false
	}

	if !checkDate(to) {
		return false
	}

	return true
}

// checkDate function check on valid date
func checkDate(array []string) bool {

	Day, err := strconv.Atoi(array[0])
	if err != nil {
		return false
	}

	if Day > 31 || Day < 1 {
		return false
	}

	Month, err := strconv.Atoi(array[1])
	if err != nil {
		return false
	}

	if Month > 12 || Month < 1 {
		return false
	}

	if Month == 4 || Month == 6 || Month == 9 || Month == 11 {
		if Day > 30 {
			return false
		}
	}

	Year, err := strconv.Atoi(array[2])
	if err != nil {
		return false
	}

	if Month == 2 && Year%4 == 0 {
		if Day > 29 {
			return false
		}
	}

	if Month == 2 && Year%4 != 0 {
		if Day > 28 {
			return false
		}
	}

	if Year > 3000 || Year < 1 {
		return false
	}
	return true
}

// CheckFirstAlbumDate function return sorted Artists and return Error in function
func CheckFirstAlbumDate(ArtistFilter []Artists, From, To string) ([]Artists, bool) {

	from := strings.Split(From, "-")
	to := strings.Split(To, "-")

	fr, err := separationArray(from)
	if !err {
		return []Artists{}, false
	}

	t, err := separationArray(to)
	if !err {
		return []Artists{}, false
	}

	if len(ArtistFilter) > 0 {
		FilterInFilter := []Artists{}
		for _, value := range ArtistFilter {

			FAD := strings.Split(value.FirstAlbum, "-")
			fad, err := separationArray(FAD)
			if !err {
				return []Artists{}, false
			}

			if comparison(fad, fr, t) {
				FilterInFilter = append(FilterInFilter, value)
			}
		}
		return FilterInFilter, true
	}

	for _, value := range ArtistsNew {

		FAD := strings.Split(value.FirstAlbum, "-")
		fad, err := separationArray(FAD)
		if !err {
			return []Artists{}, false
		}

		if comparison(fad, fr, t) {
			ArtistFilter = append(ArtistFilter, value)
		}
	}

	return ArtistFilter, true
}

// comparison function compares Dates on correspondence and return bool value
func comparison(fad, fr, t []int) bool {
	Fad := time.Month(fad[1])
	from := time.Month(fr[1])
	to := time.Month(t[1])

	FAD := time.Date(fad[2], Fad, fad[0], 0, 0, 0, 0, time.UTC)
	FROM := time.Date(fr[2], from, fr[0], 0, 0, 0, 0, time.UTC)
	TO := time.Date(t[2], to, t[0], 0, 0, 0, 0, time.UTC)

	From := FROM.Before(FAD)
	To := TO.After(FAD)
	EqualFrom := FROM.Equal(FAD)
	EqualTo := TO.Equal(FAD)

	if EqualFrom == true || EqualTo == true {
		return true
	}

	if From == true && To == true {
		return true
	}

	return false
}

// checkOn function return true if objects correspond together and return false otherwise
func checkOn(ArtistFilter []Artists, value Artists) bool {
	for _, val := range ArtistFilter {
		if val.Name == value.Name {
			return true
		}
	}
	return false
}

// separationArray function convert from array string to array int and return Error in function
func separationArray(array []string) ([]int, bool) {
	arrayInt := []int{}
	for _, value := range array {
		Int, err := strconv.Atoi(value)
		if err != nil {
			return []int{}, false
		}
		arrayInt = append(arrayInt, Int)
	}
	return arrayInt, true
}

// CheckOnLocationOfConcerts function sort by Location
func CheckOnLocationOfConcerts(ArtistFilter []Artists, Location string) []Artists {
	if len(ArtistFilter) > 0 {
		FilterInFilter := []Artists{}
		for _, value := range ArtistFilter {

			for key := range value.DatesLocations {
				Location = strings.Replace(Location, ", ", "-", 1)
				Location = strings.Replace(Location, " ", "_", -1)
				if strings.Contains(key, strings.ToLower(Location)) {
					FilterInFilter = append(FilterInFilter, value)
					break
				}
			}
		}
		return FilterInFilter
	}

	for _, value := range ArtistsNew {

		for key := range value.DatesLocations {
			Location = strings.Replace(Location, ", ", "-", 1)
			Location = strings.Replace(Location, " ", "_", -1)
			if strings.Contains(key, strings.ToLower(Location)) {
				ArtistFilter = append(ArtistFilter, value)
				break
			}
		}
	}
	return ArtistFilter
}
