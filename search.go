package pork

import (
	"encoding/json"
	"fmt"
	"github.com/akwanmaroso/devops-go/pork/nap"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type SearchResponse struct {
	Result []*SearchResult `json:"result"`
}

type SearchResult struct {
	FullName string `json:"full_name"`
}

var SearchCmd = &cobra.Command{
	Use: "search",
	Short: "search for github repository by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		if err := SearchByKeyword; err != nil {
			log.Fatalln("Search Failed:", err)
		}
	},
}

func SearchByKeyword(keyword []string) error {
	return GithubAPI().Call("search", map[string]string{
		"query": strings.Join(keyword, "+"),
	})
}

func SearchSuccess(resp *http.Response) error {
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := SearchResponse{}
	json.Unmarshal(content, &respContent)
	for _, item := range respContent.Result {
		fmt.Println(item.FullName)
	}
	return nil
}

func SearchDefaultRouter(resp *http.Response) error {
	return fmt.Errorf("status code %d", resp.StatusCode)
}

func GetSearchResource() *nap.RestResource {
	searchRouter := nap.NewRouter()
	searchRouter.DefaultRouter = SearchDefaultRouter
	searchRouter.RegisterFunc(200, SearchSuccess)
	search := nap.NewResource("/search/repositories?q={{.query}}", "GET", searchRouter)
	return search
}

