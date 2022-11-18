package manifastier_test

import (
	"encoding/json"
	"github.com/victorfernandesraton/manifastier"
	"strings"
	"testing"
)

func TestCreateManifast(t *testing.T) {
	t.Run("test simple empty manifast", func(t *testing.T) {
		result := `{"display": "fullscreen"}`
		data := manifastier.Manifest{}

		res, err := json.Marshal(data)

		if err != nil {
			t.Fatal(err)
		}
		if strings.Compare(string(res), result) == 0 {
			t.Fatalf("error when parse string, expected %v got %v", result, res)
		}
	})

}
