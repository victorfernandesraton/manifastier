package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"path"

	"github.com/spf13/cobra"
	"github.com/victorfernandesraton/manifastier/icon"
	"github.com/victorfernandesraton/manifastier/manifest"
)

var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "exec generate manifest json",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),

	RunE: func(cmd *cobra.Command, args []string) error {
		outputPath := "./manifest.json"
		iconfolder := "./public/assets/icons"
		if len(args) == 1 {
			outputPath = args[0]
		}

		fileName := path.Join(outputPath)

		iconFolderData, err := cmd.Flags().GetString("assets")
		if iconFolderData != "" {
			iconfolder = iconFolderData
		}
		displayData, err := cmd.Flags().GetString("display")
		name, err := cmd.Flags().GetString("name")
		shortName, err := cmd.Flags().GetString("short_name")
		description, err := cmd.Flags().GetString("description")
		startUrl, err := cmd.Flags().GetString("start_url")
		scope, err := cmd.Flags().GetString("scope")
		scopeUrl, err := url.Parse(scope)
		if err != nil {
			return err
		}

		if displayData == "" {
			displayData = string(manifest.Fullscreen)
		}

		icons, err := icon.ParseIconListFromDir(iconfolder)
		if err != nil {
			return err
		}

		data := manifest.Manifest{
			Icon:            icons,
			Name:            name,
			ShortName:       shortName,
			Description:     description,
			Display:         manifest.Display(displayData),
			StartURL:        startUrl,
			BackgroundColor: "#FAFAFA",
			Scope:           *scopeUrl,
		}

		res, err := json.Marshal(data)
		log.Println(string(res))
		return ioutil.WriteFile(fileName, res, 0600)

	},
}
var RootCmd = &cobra.Command{

	Use:   "manifastier",
	Short: "manifastier is a manifest.json file generator CLI",
	Long:  "manifastier is a binary independent manifest.json file generator CLI , love by Victor Raton.",
}
