package manifest_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/victorfernandesraton/manifastier/manifest"
)

func TestCreateManifast(t *testing.T) {
	t.Run("test simple empty manifast", func(t *testing.T) {
		result := `{"display": "fullscreen"}`
		data := manifest.Manifest{}

		res, err := json.Marshal(data)

		if err != nil {
			t.Fatal(err)
		}
		if strings.Compare(string(res), result) == 0 {
			t.Fatalf("error when parse string, expected %v got %v", result, res)
		}
	})

}
