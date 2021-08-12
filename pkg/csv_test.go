package coa_test

import (
	"testing"

	coa "github.com/nateinaction/coa-district-lookup/pkg"
)

func Test_ImportAddressPartial(t *testing.T) {
	addresses, err := coa.ImportAddressPartials("addresses.csv")
	if err != nil {
		t.Error(err)
	}
	if len(addresses) != 3 {
		t.Errorf("Expected 3 addresses, got %d", len(addresses))
	}
	if addresses[0].StreetNumber != "3111" {
		t.Errorf("Expected 3111, got %s", addresses[0].StreetNumber)
	}
	if addresses[1].StreetNumber != "2124" {
		t.Errorf("Expected 2124, got %s", addresses[1].StreetNumber)
	}
}
