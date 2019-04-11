package ergast

import (
	"fmt"
	"strings"
)

// CircuitDto is a container for the json data the ergast api returns
type CircuitDto struct {
	MRData struct {
		Xmlns        string `json:"xmlns"`
		Series       string `json:"series"`
		URL          string `json:"url"`
		Limit        string `json:"limit"`
		Offset       string `json:"offset"`
		Total        string `json:"total"`
		CircuitTable struct {
			Circuits []struct {
				CircuitID   string `json:"circuitId"`
				URL         string `json:"url"`
				CircuitName string `json:"circuitName"`
				Location    struct {
					Lat      string `json:"lat"`
					Long     string `json:"long"`
					Locality string `json:"locality"`
					Country  string `json:"country"`
				} `json:"Location"`
			} `json:"Circuits"`
		} `json:"CircuitTable"`
	} `json:"MRData"`
}

// GetCircuit can be filtered using one or more of the following criteria:
// /constructors/<constructorId>
// /drivers/<driverId>
// /grid/<position>
// /results/<position>
// /fastest/<rank>
// /status/<statusId>
func GetCircuit(season string, params map[string]string) (*CircuitDto, error) {
	var paramSlice []string
	var res CircuitDto
	var url string
	for key, value := range params {
		paramSlice = append(paramSlice, key, value)
	}
	joinedParams := strings.Join(paramSlice, "/")
	if params != nil {
		url = fmt.Sprintf("http://ergast.com/api/f1/%s/circuits.json", season)
	} else {
		url = fmt.Sprintf("http://ergast.com/api/f1/%s/%s/circuits.json", season, joinedParams)
	}
	err := GetObj(url, &res)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &res, err
}
