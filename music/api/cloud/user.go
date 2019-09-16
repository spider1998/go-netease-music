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

//Login
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

//ISRegister 是否已被注册
func (u UserModule) ISRegister(req types.IsRegisterRequest) (res []byte, err error) {
	/*res = readApiCache(util.StructToMapJSON(req), url.SearchMusic)
	if res != nil {
		return
	}*/
	//writeReq := util.StructToMapJSON(req)
	defaultCookies := util.SetupDefaultCookie()
	res, cookies, err := util.EAPICloudRequest(url.Host+url.ISRegisted, util.StructToMapJSON(req), defaultCookies, url.ISRegistedExt)
	if err != nil {
		return
	}
	cookies = append(cookies, defaultCookies...)
	util.WriteCookieToEnv(cookies)
	//writeApiCache(url.CellphoneLoginUrl, writeReq, res)
	return
}
