package url

const (
	Host                      = "http://music.163.com"
	LiuxUrl                   = `https://music.163.com/api/linux/forward`
	CellphoneLoginUrl         = `/weapi/login/cellphone?csrf_token=`
	RefreshLoginUrl           = `/weapi/login/token/refresh`
	Logout                    = `/weapi/login/token/logout`
	GetUserDetail             = `/weapi/v1/user/detail`
	GetUserAccountInformation = `/weapi/subcount`
	UpdateUserInformation     = `/weapi/user/profile/update`
	PlayList                  = `/weapi/user/playlist`
	PlayRecord                = `/weapi/v1/play/record`
	DjRadio                   = `/weapi/djradio/get/byuser`
	DjRadioSubed              = `/weapi/djradio/get/subed`
	GetFollows                = `/weapi/user/getfollows`
	GetFolloweds              = `/weapi/user/getfolloweds`
	GetEvent                  = `/weapi/event/get`

	GetMusicUrl      = `/weapi/song/enhance/player/url`
	SearchMusic      = `/weapi/search/get`
	GetHotSearchList = `/weapi/search/hot`
	SearchSuggest    = `/weapi/search/suggest/web`
	GetMusicLyric    = `/weapi/song/lyric?lv=-1&kv=-1&tv=-1`
	GetSongDetail    = `/weapi/v3/song/detail`
	ISRegisted       = `/eapi/cellphone/existence/check`
	ISRegistedExt    = `/api/cellphone/existence/check`
	ArtistsList      = `/weapi/v1/artist/`
	ArtistAlbums	 = `/weapi/artist/albums/`


)
