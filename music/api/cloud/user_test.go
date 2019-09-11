package api

import (
	"music/api/types"
	"music/log"
	"testing"
)

var (
	Api     *CloudAPI
	Loggers log.Logger
)

func TestUserModule_Login(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewCloudAPI(Loggers, "")
	res, err := Api.User().Login(types.CellLoginRequest{
		Phone:         "17609270263",
		Password:      "a123456",
		RememberLogin: "true",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
