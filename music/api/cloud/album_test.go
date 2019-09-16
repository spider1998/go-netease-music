package api

import (
	"encoding/json"
	"github.com/pkg/errors"
	"music/api/types"
	"music/log"
	"testing"
)

func TestAlbumModule_GetArtistAlbums(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewCloudAPI(Loggers, "")
	res, err := Api.Album().GetArtistAlbums(types.ArtistAlbumsRequest{
		Total:  true,
		Limit:  30,
		Offset: 2,
	}, "6452")
	if err != nil {
		t.Error(err)
	}
	var b types.ArtistAlbums
	err = json.Unmarshal(res, &b)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	t.Log(len(b.HotAlbums))
}
