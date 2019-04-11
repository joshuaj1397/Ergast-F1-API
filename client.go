package ergast

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	client = http.Client{
		Timeout: time.Second * 5,
	}
)

func Test() {
	circuitRes, _ := GetCircuit("2007", nil)
	fmt.Println(circuitRes)
	raceRes, _ := GetRaceResults("2010", nil)
	fmt.Println(raceRes)
}

// GetObj is a simple wrapper for GET requests
func GetObj(url string, obj interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		return getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}
	return json.Unmarshal(body, &obj)
}
