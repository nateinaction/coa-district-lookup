package coa

import "fmt"

type SearchResponse struct {
	Candidates       []Candidate `json:"candidates"`
	SpatialReference struct {
		LatestWkid int64 `json:"latestWkid"`
		Wkid       int64 `json:"wkid"`
	} `json:"spatialReference"`
}

func (s SearchResponse) HighestRankingCandidate() (Candidate, error) {
	if len(s.Candidates) == 0 {
		return Candidate{}, fmt.Errorf("no address candidates found")
	}

	var highestRanking Candidate
	for _, candidate := range s.Candidates {
		if candidate.Score > highestRanking.Score {
			highestRanking = candidate
		}
	}
	return highestRanking, nil
}
