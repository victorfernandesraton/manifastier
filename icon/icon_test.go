package icon_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/victorfernandesraton/manifastier/icon"
)

func TestParseIconToJSON(t *testing.T) {
	stub := icon.Icon{
		Size: &icon.Size{
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

func TestFindIconsByDir(t *testing.T) {
	t.Run("simple icon from dir", func(t *testing.T) {
		data, err := icon.ParseIconListFromDir("../testdata/data")
		if err != nil {
			t.Fatalf("expected nil error , got %v", err)
		}

		if data == nil {
			t.Fatal("not returning a result")
		}

		if data[0].Size.Height != 32 || data[0].Size.Width != 32 {
			t.Fatalf("not valid size expected %vx%v, got %vx%v", data[0].Size.Height, data[0].Size.Width, 32, 32)
		}
	})
	t.Run("error when directory does not exist", func(t *testing.T) {
		data, err := icon.ParseIconListFromDir("../ok")
		if data != nil {
			t.Fatalf("expected nil data, got %v", data)
		}
		if err.Error() != "open ../ok: no such file or directory" {
			t.Fatalf("expected open ../ok: no such file or directory, got %v", err)
		}
	})
}
