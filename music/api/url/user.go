package url

const (
	Host                      = "http://music.163.com"
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
	GetMusicLyric    = `/weapi/song/lyric?os=osx&id=`
	GetSongDetail    = `/weapi/v3/song/detail`
)
