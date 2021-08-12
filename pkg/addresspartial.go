package coa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AddressPartial struct {
	StreetNumber string
	StreetName   string
	Zipcode      string
}

// Search looks up address from partial data and provides a list of possible matches
func (a AddressPartial) Search() (SearchResponse, error) {
	req, err := http.NewRequest("GET", "https://www.austintexas.gov/gis/rest/Geocode/COA_Address_Locator/GeocodeServer/findAddressCandidates", nil)
	if err != nil {
		return SearchResponse{}, err
	}

	q := req.URL.Query()
	q.Add("maxLocations", "100")
	q.Add("f", "pjson")
	q.Add("Street", fmt.Sprintf("%s %s %s", a.StreetNumber, a.StreetName, a.Zipcode))
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return SearchResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SearchResponse{}, err
	}

	var searchData SearchResponse
	err = json.Unmarshal(body, &searchData)
	if err != nil {
		return SearchResponse{}, err
	}
	return searchData, nil
}
