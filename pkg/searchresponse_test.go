package coa_test

import (
	"testing"

	coa "github.com/nateinaction/coa-district-lookup/pkg"
)

func Test_HighestRankingCandidate(t *testing.T) {
	expected := coa.Candidate{
		Address: "C",
		Score:   3.0,
	}
	search := coa.SearchResponse{
		Candidates: []coa.Candidate{
			{
				Address: "A",
				Score:   1.0,
			},
			{
				Address: "B",
				Score:   2.0,
			},
			expected,
		},
	}
	if result, err := search.HighestRankingCandidate(); result != expected {
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_HighestRankingCandidate_NoResults(t *testing.T) {
	search := coa.SearchResponse{
		Candidates: []coa.Candidate{},
	}
	if _, err := search.HighestRankingCandidate(); err == nil {
		t.Errorf("Expected error but received nil")
	}
}
