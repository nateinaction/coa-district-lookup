package coa

type CouncilDistrictResponse struct {
	DisplayFieldName string `json:"displayFieldName"`
	Features         []struct {
		Attributes struct {
			CouncilDistrict int64 `json:"COUNCIL_DISTRICT"`
		} `json:"attributes"`
	} `json:"features"`
	FieldAliases struct {
		CouncilDistrict string `json:"COUNCIL_DISTRICT"`
	} `json:"fieldAliases"`
	Fields []struct {
		Alias string `json:"alias"`
		Name  string `json:"name"`
		Type  string `json:"type"`
	} `json:"fields"`
}
