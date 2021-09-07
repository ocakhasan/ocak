package weather

import (
	"os"

	"github.com/ocakhasan/ocak/cmd"
	"github.com/ocakhasan/ocak/internal/printer"

	"github.com/spf13/cobra"
)

var city string

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "hava durumu gururla sunar",
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetWeatherData(city); err != nil {
			printer.Error(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(weatherCmd)
	weatherCmd.Flags().StringVarP(&city, "city", "c", "", "City name to print the weather")
	weatherCmd.MarkFlagRequired("city")
}
