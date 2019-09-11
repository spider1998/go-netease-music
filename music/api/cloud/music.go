package api

import (
	"music/api/types"
	"music/api/url"
	"music/util"
)

type MusicModule struct {
	api *CloudAPI
}

func (m MusicModule) SearchMusic(req types.SearchParams) (res []byte, err error) {
	defaultCookies, err := util.GetCookiesForEnv()
	if err != nil {
		return
	}
	res, cookies, err := util.CloudRequest(url.Host+url.SearchMusic, util.StructToMapJSON(req), defaultCookies)
	if err != nil {
		return
	}
	cookies = append(cookies, defaultCookies...)
	util.WriteCookieToEnv(cookies)
	return
}
