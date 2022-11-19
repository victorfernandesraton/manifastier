package main

import (
	"github.com/spf13/cobra"
)

var (
	iconFolderPathArg string
	displayArg        string
	nameArg           string
	shorNameArg       string
	descriptionArg    string
	StartURLArg       string
	BackgroundColor   string
)

func main() {

	var RootCmd = &cobra.Command{
		Use:   "manifastier",
		Short: "manifastier is a manifest.json file generator CLI",
		Long:  "manifastier is a binary independent manifest.json file generator CLI , love by Victor Raton.",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	RootCmd.Flags().StringVarP(&iconFolderPathArg, "icon-folder", "i", "", "Icons folder (required)")
	RootCmd.Flags().StringVarP(&displayArg, "display", "d", "", "Display type")
	RootCmd.Flags().StringVarP(&nameArg, "name", "n", "", "Name property (required)")
	RootCmd.Flags().StringVarP(&descriptionArg, "description", "desc", "", "Start url proprety (required)")
	RootCmd.Flags().StringVarP(&shorNameArg, "short_name", "s", "", "Short name property (required)")
	RootCmd.Flags().StringVarP(&StartURLArg, "start_url", "url", "", "Start url proprety (required)")

	RootCmd.Flags().StringVarP(&BackgroundColor, "background_color", "bg", "", "Bacground color proprety (required)")

	RootCmd.Execute()
}
