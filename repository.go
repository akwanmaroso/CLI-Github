package clio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/akwanmaroso/devops-go/clio/nap"
	"github.com/spf13/cobra"
)

var (
	repoDescription string
	repoPrivate     bool
)

type RepositoryPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	Homepage    string `json:"homepage"`
}

type RepositoryResponse struct {
	FullName string `json:"full_name"`
}

var RepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "create repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must supply name repository")
		}
		if err := CreateRepository(args[0]); err != nil {
			log.Fatalln("Error when create a repository :", err)
		}
	},
}

func CreateRepository(repository string) error {
	values := strings.Split(repository, "/")
	if len(values) != 2 {
		return fmt.Errorf("repository must be in format user/repo")
	}
	payload := &RepositoryPayload{
		Name:        values[1],
		Description: repoDescription,
		Private:     repoPrivate,
		Homepage:    "https://github.com",
	}

	return GithubAPI().Call("repository", map[string]string{
		"owner":   values[0],
		"project": values[1],
	}, payload)
}

func CreateRepositorySuccess(resp *http.Response) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := RepositoryResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Printf("Create Repository: %s\n", respContent.FullName)
	return nil
}

func RepositoryDefaultRouter(resp *http.Response) error {
	return fmt.Errorf("status code %d", resp.StatusCode)
}

func GetRepositoryResource() *nap.RestResource {
	router := nap.NewRouter()
	router.RegisterFunc(201, CreateRepositorySuccess)
	router.DefaultRouter = RepositoryDefaultRouter
	resource := nap.NewResource("/user/repos", "POST", router)
	fmt.Println(resource)
	return resource
}

func init() {
	RepositoryCmd.Flags().StringVarP(&repoDescription, "description", "d", "", "description for repository")
	RepositoryCmd.Flags().BoolVarP(&repoPrivate, "private", "p", false, "Is repository private or public")
}
