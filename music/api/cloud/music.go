package api

import (
	"music/api/types"
	"music/api/url"
	"music/util"
)

type MusicModule struct {
	api *CloudAPI
}

//GetLyrics Get single lyrics
func (m MusicModule) GetLyrics(request types.GetLyricsRequest) (res []byte, err error) {
	defaultCookies, err := util.GetCookiesForEnv()
	if err != nil {
		return
	}
	req := util.StructToMapJSON(request)
	writeReq := util.StructToMapJSON(request)
	res = readApiCache(req, url.GetMusicLyric)
	if res != nil {
		return
	}
	res, cookies, err := util.LAPICloudRequest(url.Host+url.GetMusicLyric, req, defaultCookies)
	if err != nil {
		return
	}
	cookies = append(cookies, defaultCookies...)
	util.WriteCookieToEnv(cookies)
	writeApiCache(url.GetMusicLyric, writeReq, res)
	return
}

//ArtistsList Get a list of artists' popular singles（50）
func (m MusicModule) ArtistsList(id string, request types.Base) (res []byte, err error) {
	defaultCookies, err := util.GetCookiesForEnv()
	if err != nil {
		return
	}
	req := util.StructToMapJSON(request)
	writeReq := util.StructToMapJSON(request)
	res = readApiCache(req, url.ArtistsList+id)
	if res != nil {
		return
	}
	res, cookies, err := util.CloudRequest(url.Host+url.ArtistsList+id, req, defaultCookies)
	if err != nil {
		return
	}
	cookies = append(cookies, defaultCookies...)
	util.WriteCookieToEnv(cookies)
	writeApiCache(url.ArtistsList+id, writeReq, res)
	return
}

//SearchMusic Search
func (m MusicModule) SearchMusic(request types.SearchParams) (res []byte, err error) {
	defaultCookies, err := util.GetCookiesForEnv()
	if err != nil {
		return
	}
	req := util.StructToMapJSON(request)
	writeReq := util.StructToMapJSON(request)
	res = readApiCache(req, url.SearchMusic)
	if res != nil {
		return
	}
	res, cookies, err := util.CloudRequest(url.Host+url.SearchMusic, req, defaultCookies)
	if err != nil {
		return
	}
	cookies = append(cookies, defaultCookies...)
	util.WriteCookieToEnv(cookies)
	writeApiCache(url.SearchMusic, writeReq, res)
	return
}

//34 101 121 74 121 90 88 78 49 98 72 81 105 79 110 115 105 99 50 57 117 90 51 77 105 79 108 116 55 73
//34 101 121 74 121 90 88 78 49 98 72 81 105 79 110 115 105 99 50 57 117 90 51 77 105 79 108 116 55 73 109 108 107 73 106 111 120

//110 81 50 57 49 98 110 81 105 79 106 89 119 77 72 48 115 73 109 78 118 90 71 85 105 79 106 73 119 77 72 48 61 34
