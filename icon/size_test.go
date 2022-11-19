package icon_test

import (
	"testing"

	"github.com/victorfernandesraton/manifastier/icon"
)

func TestParseSizeFromByteString(t *testing.T) {
	t.Run("generate a sucessful size", func(t *testing.T) {
		input := []byte("45x50")
		stub := icon.Size{}
		err := stub.Unmarshal(input)
		if err != nil {
			t.Fatalf("expected a nil error, got %v", err)
		}
		if stub.Width != 45 {
			t.Fatalf("error in width expect %v, got %v", 45, stub.Width)
		}
		if stub.Height != 50 {
			t.Fatalf("error in height expect %v, got %v", 50, stub.Height)
		}
	})

	t.Run("error when using negative number", func(t *testing.T) {
		input := []byte("-45x50")
		stub := icon.Size{}
		err := stub.Unmarshal(input)
		if err != icon.NotValidStringAsUint {
			t.Fatalf("expected error %v, got %v", icon.NotValidStringAsUint, err)
		}
		if stub.Width != 0 {
			t.Fatalf("error in width expect %v, got %v", 0, stub.Width)
		}
		if stub.Height != 0 {
			t.Fatalf("error in height expect %v, got %v", 0, stub.Height)
		}
	})

	t.Run("error when passing more tham one valuer", func(t *testing.T) {
		input := []byte("45x50x43")
		stub := icon.Size{}
		err := stub.Unmarshal(input)
		if err != icon.NotHaveAValidSizes {
			t.Fatalf("expected error %v, got %v", icon.NotHaveAValidSizes, err)
		}
		if stub.Width != 0 {
			t.Fatalf("error in width expect %v, got %v", 0, stub.Width)
		}
		if stub.Height != 0 {
			t.Fatalf("error in height expect %v, got %v", 0, stub.Height)
		}
	})

	t.Run("error when not have a separator", func(t *testing.T) {
		input := []byte("45")
		stub := icon.Size{}
		err := stub.Unmarshal(input)
		if err != icon.NotSeparatorError {
			t.Fatalf("expected error %v, got %v", icon.NotSeparatorError, err)
		}
		if stub.Width != 0 {
			t.Fatalf("error in width expect %v, got %v", 0, stub.Width)
		}
		if stub.Height != 0 {
			t.Fatalf("error in height expect %v, got %v", 0, stub.Height)
		}
	})
}
