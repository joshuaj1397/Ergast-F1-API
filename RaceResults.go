package ergast

import (
	"fmt"
	"strings"
)

// RaceResultsDto is a container for the json data the ergast api returns
type RaceResultsDto struct {
	MRData struct {
		Xmlns     string `json:"xmlns"`
		Series    string `json:"series"`
		URL       string `json:"url"`
		Limit     string `json:"limit"`
		Offset    string `json:"offset"`
		Total     string `json:"total"`
		RaceTable struct {
			Season string `json:"season"`
			Round  string `json:"round"`
			Races  []struct {
				Season   string `json:"season"`
				Round    string `json:"round"`
				URL      string `json:"url"`
				RaceName string `json:"raceName"`
				Circuit  struct {
					CircuitID   string `json:"circuitId"`
					URL         string `json:"url"`
					CircuitName string `json:"circuitName"`
					Location    struct {
						Lat      string `json:"lat"`
						Long     string `json:"long"`
						Locality string `json:"locality"`
						Country  string `json:"country"`
					} `json:"Location"`
				} `json:"Circuit"`
				Date    string `json:"date"`
				Time    string `json:"time"`
				Results []struct {
					Number       string `json:"number"`
					Position     string `json:"position"`
					PositionText string `json:"positionText"`
					Points       string `json:"points"`
					Driver       struct {
						DriverID        string `json:"driverId"`
						PermanentNumber string `json:"permanentNumber"`
						Code            string `json:"code"`
						URL             string `json:"url"`
						GivenName       string `json:"givenName"`
						FamilyName      string `json:"familyName"`
						DateOfBirth     string `json:"dateOfBirth"`
						Nationality     string `json:"nationality"`
					} `json:"Driver,omitempty"`
					Constructor struct {
						ConstructorID string `json:"constructorId"`
						URL           string `json:"url"`
						Name          string `json:"name"`
						Nationality   string `json:"nationality"`
					} `json:"Constructor"`
					Grid   string `json:"grid"`
					Laps   string `json:"laps"`
					Status string `json:"status"`
					Time   struct {
						Millis string `json:"millis"`
						Time   string `json:"time"`
					} `json:"Time,omitempty"`
					FastestLap struct {
						Rank string `json:"rank"`
						Lap  string `json:"lap"`
						Time struct {
							Time string `json:"time"`
						} `json:"Time"`
						AverageSpeed struct {
							Units string `json:"units"`
							Speed string `json:"speed"`
						} `json:"AverageSpeed"`
					} `json:"FastestLap,omitempty"`
				} `json:"Results"`
			} `json:"Races"`
		} `json:"RaceTable"`
	} `json:"MRData"`
}

// GetRaceResults can be filtered using one or more of the following criteria:
//
// /circuits/<circuitId>
// /constructors/<constructorId>
// /drivers/<driverId>
// /grid/<position>
// /fastest/<rank>
// /status/<statusId>
func GetRaceResults(season string, params map[string]string) (*RaceResultsDto, error) {
	var paramSlice []string
	var res RaceResultsDto
	var url string
	for key, value := range params {
		paramSlice = append(paramSlice, key, value)
	}
	joinedParams := strings.Join(paramSlice, "/")
	if params != nil {
		url = fmt.Sprintf("http://ergast.com/api/f1/%s/%s/results.json", season, joinedParams)
	} else {
		url = fmt.Sprintf("http://ergast.com/api/f1/%s/results.json", season)
	}
	err := GetObj(url, &res)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &res, err
}
