package pork

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/akwanmaroso/devops-go/pork/nap"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ReadmeResponse struct {
	Content string `json:"content"`
}

var DocsCmd = &cobra.Command{
	Use: "docs",
	Short: "read the documentation for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must supply repository argument")
		}
		if err := GetRepositoryReadme(args[0]); err != nil {
			log.Fatalln("Failed to get docs: ", err)
		}
	},
}

func GetRepositoryReadme(repository string) error {
	values := strings.Split(repository, "/")
	return GithubAPI().Call("docs", map[string]string{
		"owner": values[0],
		"project": values[1],
	}, nil)
}

func ReadmeSuccess(resp *http.Response) error {
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respContent := ReadmeResponse{}
	json.Unmarshal(content, &respContent)
	buff, err := base64.StdEncoding.DecodeString(respContent.Content)
	if err != nil {
		return err
	}
	fmt.Println(string(buff))
	return nil
}

func ReadmeDefaultRouter(resp *http.Response) error {
	return fmt.Errorf("status code %d", resp.StatusCode)
}

func GetReadmeResource() *nap.RestResource {
	router := nap.NewRouter()
	router.RegisterFunc(200, ReadmeSuccess)
	router.DefaultRouter = ReadmeDefaultRouter
	resource := nap.NewResource("/repos/{{.owner}}/{{.project}}/readme", "GET", router)
	return resource
}
