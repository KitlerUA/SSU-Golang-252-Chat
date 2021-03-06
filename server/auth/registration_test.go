package auth

import (
	"testing"

	"fmt"

	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)

func TestRegisterNewUser(t *testing.T) {

	test_user := messageService.Authentification{
		UserName: "fklds222d",
		Password: "vasya22sss",
		NickName: "voron",
	}
	user, tok, err := RegisterNewUser(&test_user)
	fmt.Println(user, tok, err)
}
