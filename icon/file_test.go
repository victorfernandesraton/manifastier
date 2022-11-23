package icon_test

import (
	"testing"

	"github.com/victorfernandesraton/manifastier/icon"
)

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
	t.Run("empty list of icons", func(t *testing.T) {
		data, err := icon.ParseIconListFromDir("../testdata/empty")
		if data != nil {
			t.Fatalf("expected nil data, got %v", data)
		}
		if err != nil {
			t.Fatalf("expected nilerror, got %v", err)
		}
	})
}
