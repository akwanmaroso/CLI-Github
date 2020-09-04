package clio

import (
	"github.com/akwanmaroso/devops-go/clio/nap"
	"github.com/spf13/viper"
	"testing"
)

func TestForkRepository(t *testing.T) {
	token := viper.GetString("token")
	GithubAPI().SetAuth(nap.NewAuthToken(token))
	if err := ForkRepository("avelino/awesome-go"); err != nil {
		t.Fail()
	}
}