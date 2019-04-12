package ergast

import (
	"log"
	"testing"
)

func TestGetCircuitWrapper(t *testing.T) {
	circuitRes, err := GetCircuit("2007", nil)

	if err != nil {
		t.Error(err)
	}

	if circuitRes == nil {
		t.Error("No results. Most likely a bad request")
	}

	log.Println(circuitRes)
}
