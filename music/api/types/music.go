package types

type Base struct{}

// @Title SearchMusic
// @Description Search music
// @Params keywords   	 query    string    true    "search criteria keywords"
// @Params type		   	 query    int	    false   "search type, 默认为 1 即单曲 , 取值意义 : 1: 单曲, 10: 专辑, 100: 歌手, 1000: 歌单, 1002: 用户, 1004: MV, 1006: 歌词, 1009: 电台, 1014: 视频"
// @Params offset   	 query    int       false   "items offset"
// @Params limit   	 	 query    int       false   "items limit"
// @Success 200 {string}
type SearchParams struct {
	Keywords string `json:"s"`
	Type     int    `json:"type"`
	Limit    int    `json:"limit" default:"30"`
	Offset   int    `json:"offset" default:"0"`
}

type Artist struct {
	Name        string `json:"name"`
	TopicPerson int    `json:"topicPerson"`
	AlbumSize   int    `json:"albumSize"`
	MusicSize   int    `json:"musicSize"`
	MVSize      int    `json:"mvSize"`
}

type HotSong struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	EQ   string `json:"eq"`
	AR   []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"ar"`
	AL struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"al"`
}

type Artists struct {
	Artist   Artist    `json:"artist"`
	HotSongs []HotSong `json:"hotSongs"`
	More     bool      `json:"more"`
	Code     int       `json:"code"`
}

type GetLyricsRequest struct {
	ID int `json:"id"`
}

type Lyrics struct {
	LRC struct {
		Version int    `json:"version"`
		Lyric   string `json:"lyric"`
	} `json:"lrc"`
	Code int `json:"code"`
}
