package manifastier

import (
	"errors"

	"github.com/victorfernandesraton/manifastier/icon"
)

var (
	NotValidWIthoutIcon = errors.New("not valid because needed one or more icons")
)

type Manifest struct {
	Name            string       `json:"name"`
	ShortName       string       `json:"short_name"`
	StartURL        string       `json:"start_url"`
	Display         Display      `json:"display"`
	BackgroundColor string       `json:"background_color"`
	Description     string       `json:"description"`
	Icon            []*icon.Icon `json:"icon"`
}

func New(name string, short_name string, startURL string, display Display, backgroundColor string, description string, iconList []string (*Manifest, error) {
	if len(iconList) < 1 {
		return nil, NotValidWIthoutIcon
	}


}
