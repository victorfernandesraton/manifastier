package manifest

import (
	"errors"
	"net/url"

	"github.com/victorfernandesraton/manifastier/icon"
)

var (
	NotValidWIthoutIcon = errors.New("not valid because needed one or more icons")
)

type Display string

const (
	Fullscreen Display = "fullscreen"
	Standalone         = "standalone"
	MinimalUi          = "minimal-ui"
	Browser            = "browser"
)

type Manifest struct {
	Name            string       `json:"name"`
	ShortName       string       `json:"short_name"`
	StartURL        string       `json:"start_url"`
	Display         Display      `json:"display"`
	BackgroundColor string       `json:"background_color"`
	Description     string       `json:"description"`
	Icon            []*icon.Icon `json:"icon"`
	Scope           url.URL      `json:"scope"`
}
