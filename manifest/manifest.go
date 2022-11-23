package manifest

import (
	"errors"
	"net/url"

	"github.com/victorfernandesraton/manifastier/display"
	"github.com/victorfernandesraton/manifastier/icon"
)

var (
	NotValidWIthoutIcon = errors.New("not valid because needed one or more icons")
)

type Manifest struct {
	Name            string          `json:"name"`
	ShortName       string          `json:"short_name"`
	StartURL        string          `json:"start_url"`
	Display         display.Display `json:"display"`
	BackgroundColor string          `json:"background_color"`
	Description     string          `json:"description"`
	Icon            []*icon.Icon    `json:"icon"`
	Scope           url.URL         `json:"scope"`
}
