package icon

import (
	"encoding/json"
	"fmt"
)

type Icon struct {
	Size *Size  `json:"size"`
	Src  string `json:"src"`
}

type alias struct {
	Size string `json:"size"`
	Src  string `json:"src"`
}

func (i *Icon) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(&alias{
		Src:  i.Src,
		Size: fmt.Sprintf("%vx%v", i.Size.Width, i.Size.Height),
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}
