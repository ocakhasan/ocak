package imdb

import (
	"os"

	"github.com/ocakhasan/ocak/cmd"
	"github.com/ocakhasan/ocak/internal/printer"
	"github.com/spf13/cobra"
)

var movie string

var imdbCmd = &cobra.Command{
	Use:   "imdb",
	Short: "search movies through imdb",
	Run: func(cmd *cobra.Command, args []string) {
		if err := findMovie(movie); err != nil {
			printer.Error(os.Stdout, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(imdbCmd)
	imdbCmd.Flags().StringVarP(&movie, "movie", "m", "Godfather", "Movie name")
	imdbCmd.MarkFlagRequired("movie")
}
