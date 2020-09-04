package clio

import (
	"github.com/akwanmaroso/devops-go/clio/nap"
	"github.com/spf13/viper"
	"testing"
)

func TestGetRepositoryReadme(t *testing.T) {
	token := viper.GetString("token")
	GithubAPI().SetAuth(nap.NewAuthToken(token))
	if err := GetRepositoryReadme("akwanmaroso/clio"); err != nil {
		t.Fail()
	}
}
