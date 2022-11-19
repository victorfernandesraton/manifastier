package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/spf13/cobra"
	"github.com/victorfernandesraton/manifastier/display"
	"github.com/victorfernandesraton/manifastier/icon"
	"github.com/victorfernandesraton/manifastier/manifest"
)

var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "exec generate manifest json",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),

	RunE: func(cmd *cobra.Command, args []string) error {
		outputPath := "./manifest.json"
		if len(args) == 1 {
			outputPath = args[0]
		}

		fileName := path.Join(outputPath)

		iconFolder, err := cmd.Flags().GetString("icon-folder")
		displayData, err := cmd.Flags().GetString("display")
		name, err := cmd.Flags().GetString("name")
		shortName, err := cmd.Flags().GetString("short_name")
		description, err := cmd.Flags().GetString("description")
		startUrl, err := cmd.Flags().GetString("start_url")
		if err != nil {
			return err
		}

		if displayData == "" {
			displayData = string(display.Fullscreen)
		}

		icons, err := icon.ParseIconListFromDir(iconFolder)
		if err != nil {
			return err
		}

		data := manifest.Manifest{
			Icon:            icons,
			Name:            name,
			ShortName:       shortName,
			Description:     description,
			Display:         display.Display(displayData),
			StartURL:        startUrl,
			BackgroundColor: "#FAFAFA",
		}

		res, err := json.Marshal(data)
		fmt.Println(string(res))
		return ioutil.WriteFile(fileName, res, 0600)

	},
}
var RootCmd = &cobra.Command{

	Use:   "manifastier",
	Short: "manifastier is a manifest.json file generator CLI",
	Long:  "manifastier is a binary independent manifest.json file generator CLI , love by Victor Raton.",
}
