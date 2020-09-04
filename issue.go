package clio

import (
	"encoding/json"
	"fmt"
	"github.com/akwanmaroso/devops-go/clio/nap"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	issueTitle string
	issueMessage string
)

type IssuePayload struct {
	Title string `json:"title"`
	Body string `json:"body"`
}

type IssueResponse struct {
	URL string `json:"html_url"`
}

var IssueCmd = &cobra.Command{
	Use: "issue",
	Short: "create issue",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must supply a repository")
		}
		if err := CreateIssue(args[0]); err != nil {
			log.Fatalln("Failed to create issue: ", err)
		}
	},
}

func CreateIssue(repository string) error {
	values := strings.Split(repository, "/")
	if len(values) != 2 {
		return fmt.Errorf("Repository must be in format owner/project")
	}
	payload := &IssuePayload{
		Title: issueTitle,
		Body: issueMessage,
	}

	return GithubAPI().Call("issue", map[string]string{
		"owner": values[0],
		"project": values[1],
	}, payload)
}

func IssueSuccess(resp *http.Response) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := IssueResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Printf("Create Issue: %s\n", respContent.URL)
	return nil
}

func IssueDefaultRouter(resp *http.Response) error {
	return fmt.Errorf("status code %d", resp.StatusCode)
}

func GetIssueResource() *nap.RestResource {
	router := nap.NewRouter()
	router.RegisterFunc(201, IssueSuccess)
	router.DefaultRouter = IssueDefaultRouter
	resource := nap.NewResource("/repos/{{.owner}}/{{.project}}/issues", "POST", router)
	return resource
}

func init() {
	IssueCmd.Flags().StringVarP(&issueTitle, "title", "t", "Basic create issue", "title issue")
	IssueCmd.Flags().StringVarP(&issueMessage, "message", "m", "", "message issue")
}

