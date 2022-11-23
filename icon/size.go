package icon

import (
	"errors"
	"log"
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

func UnmarshalSize(data []byte) (*Size, error) {
	dataString := string(data)
	var result Size
	if !strings.Contains(dataString, separator) {
		return nil, NotSeparatorError
	}

	listValues := strings.Split(dataString, separator)

	if len(listValues) != 2 {
		return nil, NotHaveAValidSizes
	}

	for i, value := range listValues {
		res, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nil, NotValidStringAsUint
		}
		if i == 0 {
			result.Width = uint64(res)
		} else {
			result.Height = uint64(res)
		}
	}
	log.Println(result)
	return &result, nil
}
