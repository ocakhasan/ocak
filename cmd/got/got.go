package got

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
	Use:   "got",
	Short: "game of thrones demeçleri",
	Run: func(cmd *cobra.Command, args []string) {
		err := randomQuote()
		if err != nil {
			printer.Error(os.Stderr, "%v\n", err)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(imdbCmd)
}

const URL string = "https://game-of-thrones-quotes.herokuapp.com/v1/random"

type Quote struct {
	Sentence  string    `json:"sentence"`
	Character Character `json:"character"`
}

type Character struct {
	Name string `json:"name"`
}

func (q Quote) String() string {
	return fmt.Sprintf(`%v demiş ki: %v`, q.Character.Name, q.Sentence)
}

func randomQuote() error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code : %v\n for chuck norris joke", resp.StatusCode)
	}

	var quote Quote
	if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
		return err
	}
	printer.Success(os.Stdout, "%v\n", quote)
	return nil
}