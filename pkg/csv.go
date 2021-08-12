package coa

import (
	"encoding/csv"
	"io"
	"os"
)

func ImportAddressPartials(filename string) ([]AddressPartial, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var addressPartials []AddressPartial
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		addressPartials = append(addressPartials, AddressPartial{
			StreetNumber: record[0],
			StreetName:   record[1],
			Zipcode:      record[2],
		})
	}
	return addressPartials, nil
}
