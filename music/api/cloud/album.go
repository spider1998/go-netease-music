package api

import (
	"music/api/types"
	"music/api/url"
	"music/util"
)

type AlbumModule struct {
	api *CloudAPI
}

func (a AlbumModule) GetArtistAlbums(request types.ArtistAlbumsRequest, id string) (res []byte, err error) {
	defaultCookies, err := util.GetCookiesForEnv()
	if err != nil {
		return
	}
	req := util.StructToMapJSON(request)
	writeReq := util.StructToMapJSON(request)
	res = readApiCache(req, url.ArtistAlbums+id)
	if res != nil {
		return
	}
	res, cookies, err := util.CloudRequest(url.Host+url.ArtistAlbums+id, req, defaultCookies)
	if err != nil {
		return
	}
	cookies = append(cookies, defaultCookies...)
	util.WriteCookieToEnv(cookies)
	writeApiCache(url.ArtistAlbums+id, writeReq, res)
	return
}
