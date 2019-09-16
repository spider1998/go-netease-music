package api

import (
	"encoding/json"
	"github.com/pkg/errors"
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

func TestUserModule_ISRegister(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewCloudAPI(Loggers, "")
	res, err := Api.User().ISRegister(types.IsRegisterRequest{
		Cellphone: "17609270263",
	})
	if err != nil {
		t.Error(err)
	}
	var b map[string]interface{}
	err = json.Unmarshal(res, &b)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	t.Log(b)
}

///api/cellphone/existence/check-36cd479b6b5-{"cellphone":"17609270264","header":
// {"appver":"6.1.1","versioncode":"140","buildver":"1568623008","resolution":"1920x1080",
// "__csrf":"","os":"android","requestId":"1568623008396_0176"}}-36cd479b6b5-e863dddd800098f750fe174832373a86

///eapi/cellphone/existence/check-36cd479b6b5-{"cellphone":"17609270263","header":
// {"__csrf":"","appver":"6.1.1","buildver":"1568622998","mobilename":null,"os":"android",
// "requestId":"1568622998369_0169","resolution":"1920x1080","versioncode":"140"}}
// -36cd479b6b5-d7ce1e1003f1a17be788734cadfa65e4