package icon

import (
	"errors"
	"strconv"
	"strings"
)

const separator = "x"

var (
	NotSeparatorError    = errors.New("not have a x separator")
	NotHaveAValidSizes   = errors.New("not have a valid 2 sizes params")
	NotValidStringAsUint = errors.New("not a valid string as uint64")
)

type Size struct {
	Width, Height uint64
}

func (s *Size) Unmarshal(data []byte) error {
	dataString := string(data)
	if !strings.Contains(dataString, separator) {
		return NotSeparatorError
	}

	listValues := strings.Split(dataString, separator)

	if len(listValues) != 2 {
		return NotHaveAValidSizes
	}

	for i, value := range listValues {
		res, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return NotValidStringAsUint
		}
		if i == 0 {
			s.Width = uint64(res)
		} else {
			s.Height = uint64(res)
		}
	}
	return nil
}
