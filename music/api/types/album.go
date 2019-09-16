package types

type ArtistAlbumsRequest struct {
	Limit  int  `json:"limit" default:"30"`
	Offset int  `json:"offset" default:"0"`
	Total  bool `json:"total" default:"true"`
}

type ArtistAlbums struct {
	Artist struct {
		Name      string `json:"name"`
		ID        int    `json:"id"`
		MusicSize int    `json:"musicSize"`
	} `json:"artist"`
	HotAlbums []struct {
		Company string `json:"company"`
		SubType string `json:"subType"`
		Name string `json:"name"`
		ID int `json:"id"`
		Type string `json:"type"`
		Size int `json:"size"`
	} `json:"hotAlbums"`
	More bool `json:"more"`
	Code int  `json:"code"`
}
