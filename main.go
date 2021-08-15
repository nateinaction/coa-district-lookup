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
	districtCache := map[string]district{}
	for _, addressPartial := range addressPartials {
		// Skip if address is a duplicate
		cacheKey := fmt.Sprintf("%s%s%s", addressPartial.StreetNumber, addressPartial.StreetName, addressPartial.Zipcode)
		if _, ok := districtCache[cacheKey]; ok {
			print(&addressPartial, districtCache[cacheKey])
			continue
		}

		addressSearch, err := addressPartial.Search()
		if err != nil {
			districtCache[cacheKey] = district{
				err: err,
			}
			print(&addressPartial, districtCache[cacheKey])
			continue
		}

		highestRankingAddress, err := addressSearch.HighestRankingCandidate()
		if err != nil {
			districtCache[cacheKey] = district{
				err: err,
			}
			print(&addressPartial, districtCache[cacheKey])
			continue
		}

		districtResp, err := highestRankingAddress.GetDistrict()
		if err != nil {
			districtCache[cacheKey] = district{
				err: err,
			}
			print(&addressPartial, districtCache[cacheKey])
			continue
		}
		districtCache[cacheKey] = district{
			district: districtResp,
		}
		print(&addressPartial, districtCache[cacheKey])
	}
}

type district struct {
	district int64
	err      error
}

func print(addressPartial *coa.AddressPartial, district district) {
	var message string
	if district.err != nil {
		message = fmt.Sprintf("%v", district.err)
	} else {
		message = fmt.Sprintf("%d", district.district)
	}

	fmt.Printf(
		"%s,%s,%s,%s\n",
		addressPartial.StreetNumber,
		addressPartial.StreetName,
		addressPartial.Zipcode,
		message,
	)
}
