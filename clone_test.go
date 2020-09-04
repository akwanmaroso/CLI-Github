package clio

import (
	"github.com/akwanmaroso/devops-go/clio/nap"
	"github.com/spf13/viper"
	"testing"
)

func TestCloneRepository(t *testing.T) {
	token := viper.GetString("token")
	GithubAPI().SetAuth(nap.NewAuthToken(token))
	if err := CloneRepository("avelino/awesome-go", "master", false); err != nil {
		t.Fail()
	}
}