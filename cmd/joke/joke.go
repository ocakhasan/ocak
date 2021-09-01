package joke

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ocakhasan/ocak/cmd"
	"github.com/ocakhasan/ocak/internal/printer"
	"github.com/spf13/cobra"
)

var imdbCmd = &cobra.Command{
	Use:   "joke",
	Short: "yap bakayım şakanı",
	Run: func(cmd *cobra.Command, args []string) {
		err := randomJoke()
		if err != nil {
			printer.Error(os.Stderr, "%v\n", err)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(imdbCmd)
}

const URL string = "https://api.chucknorris.io/jokes/random"

type Joke struct {
	Value string `json:"value"`
}

func randomJoke() error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code : %v\n for chuck norris joke", resp.StatusCode)
	}

	var joke Joke
	if err := json.NewDecoder(resp.Body).Decode(&joke); err != nil {
		return err
	}
	printer.Success(os.Stdout, "joke: %v\n", joke.Value)
	return nil
}
