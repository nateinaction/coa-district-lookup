package coa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Candidate struct {
	Address    string   `json:"address"`
	Attributes struct{} `json:"attributes"`
	Location   struct {
		X json.Number `json:"x"`
		Y json.Number `json:"y"`
	} `json:"location"`
	Score float64 `json:"score"`
}

func (c Candidate) GetDistrict() (int64, error) {
	req, err := http.NewRequest("GET", "https://www.austintexas.gov/gis/rest/Shared/CouncilDistrictsFill/MapServer/0/query", nil)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	q.Add("geometryType", "esriGeometryPoint")
	q.Add("spatialRel", "esriSpatialRelIntersects")
	q.Add("outFields", "COUNCIL_DISTRICT")
	q.Add("returnGeometry", "false")
	q.Add("returnIdsOnly", "false")
	q.Add("returnCountOnly", "false")
	q.Add("returnZ", "false")
	q.Add("returnM", "false")
	q.Add("returnDistinctValues", "false")
	q.Add("f", "pjson")
	q.Add("geometry", fmt.Sprintf("%s,%s", c.Location.X, c.Location.Y))
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var districtData CouncilDistrictResponse
	err = json.Unmarshal(body, &districtData)
	if err != nil {
		return 0, err
	}
	return districtData.Features[0].Attributes.CouncilDistrict, nil
}
