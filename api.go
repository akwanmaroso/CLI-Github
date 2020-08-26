package pork

import (
	"github.com/akwanmaroso/devops-go/pork/nap"
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
	}

	return api
}