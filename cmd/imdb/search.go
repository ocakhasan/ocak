package imdb

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/ocakhasan/ocak/cmd"
	"github.com/ocakhasan/ocak/internal/printer"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

var resultCount int64 = 10

func findMovie(movie string) error {
	keywords := strings.Join(strings.Split(movie, " "), "+")

	cse, err := createSearchService()
	if err != nil {
		return err
	}

	config := cmd.LoadConfig()
	call := cse.List().Q(keywords).Cx(config.EngineID).Num(resultCount)
	resp, err := call.Do()
	if err != nil {
		concreteErr := err.(*googleapi.Error)
		if concreteErr.Code == 403 && concreteErr.Message == "Daily Limit Exceeded" {
			return fmt.Errorf("quota exceeded")
		}
		return fmt.Errorf("error making image search API call for the given criteria: %v Err: %v", keywords, err)
	}

	if len(resp.Items) == 0 {
		return fmt.Errorf("could not find any image based for the given criteria: %v", keywords)
	}

	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	for _, url := range resp.Items {
		if strings.Contains(url.FormattedUrl, "imdb.com/title/tt") {
			line := fmt.Sprintf("%s\t%s", printer.Yellow(url.Title), printer.Cyan(url.FormattedUrl))
			fmt.Fprintln(w, line)
		}
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil

}

func createSearchService() (*customsearch.CseService, error) {
	config := cmd.LoadConfig()
	apiKey := config.CustomSearch
	searchService, err := customsearch.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("error creating customsearch client: %v", err)
	}

	cse := customsearch.NewCseService(searchService)
	return cse, nil
}
