package api

import (
	"crypto/md5"
	"fmt"
	"music/api/types"
	"music/api/url"
	"music/util"
)

type UserModule struct {
	api *CloudAPI
}

func (u UserModule) Login(req types.CellLoginRequest) (res interface{}, err error) {
	defaultCookies := util.SetupDefaultCookie()
	data := []byte(req.Password)
	has := md5.Sum(data)
	req.Password = fmt.Sprintf("%x", has)

	res, cookies, err := util.CloudRequest(url.Host+url.CellphoneLoginUrl, util.StructToMapJSON(req), defaultCookies)
	if err != nil {
		return
	}
	cookies = append(cookies, defaultCookies...)
	util.WriteCookieToEnv(cookies)
	return
}
