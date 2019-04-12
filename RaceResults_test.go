package ergast

import (
	"log"
	"testing"
)

func TestRaceResultsWrapper(t *testing.T) {
	raceResults, err := GetRaceResults("2010", nil)

	if err != nil {
		t.Error(err)
	}

	if raceResults == nil {
		t.Error("No results. Most likely a bad request")
	}

	log.Println(raceResults)
}
