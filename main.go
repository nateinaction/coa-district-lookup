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
