package icon_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/victorfernandesraton/manifastier/icon"
)

func TestParseIconToJSON(t *testing.T) {
	stub := icon.Icon{
		Size: icon.Size{
			Width:  50,
			Height: 65,
		},
		Src: "path-to-test",
	}

	jsonString := `{"size":"50x65","src":"path-to-test"}`

	res, err := json.Marshal(&stub)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if strings.Compare(string(res), jsonString) != 0 {
		t.Fatalf("expected %v, got %v", jsonString, string(res))
	}

}
