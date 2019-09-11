package types

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
	Offset   int    `json:"offset"`
}
