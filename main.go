package main

import (
	"fmt"
	"log"
	"os"

	coa "github.com/nateinaction/coa-district-lookup/pkg"
)

func main() {
	// Import CSV data
	// Expected format:
	// StreetNumber,StreetName,Zipcode
	filename := os.Args[1]
	addressPartials, err := coa.ImportAddressPartials(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Lookup district for each address and place in cache
	districtCache := map[string]int64{}
	for _, addressPartial := range addressPartials {
		// Skip if address is a duplicate
		cacheKey := fmt.Sprintf("%s%s%s", addressPartial.StreetNumber, addressPartial.StreetName, addressPartial.Zipcode)
		if _, ok := districtCache[cacheKey]; ok {
			fmt.Printf(
				"%s, %s, %s, %d\n",
				addressPartial.StreetNumber,
				addressPartial.StreetName,
				addressPartial.Zipcode,
				districtCache[cacheKey],
			)
			continue
		}

		addressSearch, err := addressPartial.Search()
		if err != nil {
			log.Fatal(err)
		}

		highestRankingAddress, err := addressSearch.HighestRankingCandidate()
		if err != nil {
			log.Fatal(err)
		}

		district, err := highestRankingAddress.GetDistrict()
		if err != nil {
			log.Fatal(err)
		}
		districtCache[cacheKey] = district
		fmt.Printf(
			"%s, %s, %s, %d\n",
			addressPartial.StreetNumber,
			addressPartial.StreetName,
			addressPartial.Zipcode,
			districtCache[cacheKey],
		)
	}
}

// Full address lookup
// GET https://www.austintexas.gov/gis/rest/Shared/Property/MapServer/0/query?f=json&where=UPPER(FULL_STREET_NAME)%20LIKE%20UPPER(%273111%20PARKER%20L%25%27)&returnGeometry=true&spatialRel=esriSpatialRelIntersects&outFields=PLACE_ID%2CFULL_STREET_NAME
// f: json
// where: UPPER(FULL_STREET_NAME) LIKE UPPER('3111%')
// returnGeometry: true
// spatialRel: esriSpatialRelIntersects
// outFields: PLACE_ID,FULL_STREET_NAME

// District lookup
// GET  https://www.austintexas.gov/Geocortex/Essentials/External/REST/sites/Council_Districts/map/mapservices/12/layers/0/datalinks/AddressDetails/link?f=json&pff_PLACE_ID=995379&dojo.preventCache=1628729368683

// Full address lookup
// OR https://www.austintexas.gov/government
// GET https://www.austintexas.gov/gis/rest/Geocode/COA_Address_Locator/GeocodeServer/findAddressCandidates?outFields=&maxLocations=100&outSR=&searchExtent=&f=pjson&Street=3111%20Parker

// District lookup
// https://www.austintexas.gov/gis/rest/Shared/CouncilDistrictsFill/MapServer/0/query?where=&text=&objectIds=&time=&geometryType=esriGeometryPoint&inSR=&spatialRel=esriSpatialRelIntersects&relationParam=&outFields=COUNCIL_DISTRICT&returnGeometry=false&maxAllowableOffset=&geometryPrecision=&outSR=&returnIdsOnly=false&returnCountOnly=false&orderByFields=&groupByFieldsForStatistics=&outStatistics=&returnZ=false&returnM=false&gdbVersion=&returnDistinctValues=false&f=pjson&geometry=3114950.749263479,10054578.00033144
