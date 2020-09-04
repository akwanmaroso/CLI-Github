package clio

import (
	"github.com/akwanmaroso/devops-go/clio/nap"
	"github.com/spf13/viper"
)

var api *nap.API

func GithubAPI() *nap.API {
	if api == nil {
		api = nap.NewAPI("https://api.github.com")
		token := viper.GetString("token")
		api.SetAuth(nap.NewAuthToken(token))
		api.AddResource("fork", GetForkResource())
		api.AddResource("search", GetSearchResource())
		api.AddResource("docs", GetReadmeResource())
		api.AddResource("pullrequest", GetPullRequestResource())
		api.AddResource("issue", GetIssueResource())
	}

	return api
}
