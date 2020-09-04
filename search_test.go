package clio

import (
	"github.com/akwanmaroso/devops-go/clio/nap"
	"github.com/spf13/viper"
	"testing"
)

func TestSearchByKeyword(t *testing.T) {
	token := viper.GetString("token")
	GithubAPI().SetAuth(nap.NewAuthToken(token))
	if err := SearchByKeyword([]string{"topic:go"}); err != nil {
		t.Fail()
	}
}
