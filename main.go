package main

import (
	"log"

	"github.com/victorfernandesraton/manifastier/cmd"
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

func init() {
	cmd.RootCmd.AddCommand(cmd.ExecCmd)
}

func main() {

	cmd.ExecCmd.Flags().StringVarP(&iconFolderPathArg, "icon-folder", "i", "", "Icons folder (required)")
	cmd.ExecCmd.Flags().StringVarP(&displayArg, "display", "d", "", "Display type")
	cmd.ExecCmd.Flags().StringVarP(&nameArg, "name", "n", "", "Name property (required)")
	cmd.ExecCmd.Flags().StringVar(&descriptionArg, "description", "", "Start url proprety (required)")
	cmd.ExecCmd.Flags().StringVarP(&shorNameArg, "short_name", "s", "", "Short name property (required)")
	cmd.ExecCmd.Flags().StringVarP(&StartURLArg, "start_url", "u", "", "Start url proprety (required)")
	cmd.ExecCmd.Flags().StringVarP(&BackgroundColor, "background_color", "b", "", "Bacground color proprety (required)")

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
