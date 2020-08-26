package nap

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/labstack/gommon/log"
)

type AuthToken struct {
	Token string
}

type AuthBasic struct {
	Username string
	Password string
}

type Authentication interface {
	AuthorizationHeader() string
}

func NewAuthToken(token string) *AuthToken {
	return &AuthToken{Token: token}
}

func NewAuthBasic(username string, password string) *AuthBasic {
	return &AuthBasic{Username: username, Password: password}
}

func (a *AuthToken) AuthorizationHeader() string {
	return fmt.Sprintf("token %s", a.Token)
}

func (a *AuthBasic) AuthorizationHeader() string {
	buffer := &bytes.Buffer{}
	enc := base64.NewEncoder(base64.URLEncoding, buffer)
	encContent := fmt.Sprintf("%s:%s", a.Username, a.Password)
	enc.Write([]byte(encContent))
	enc.Close()

	content, err := ioutil.ReadAll(buffer)
	if err != nil {
		log.Fatalf("Read failed: %v\n", err)
	}

	return fmt.Sprintf("basic %v", string(content))
}
