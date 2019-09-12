package api

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"music/api/types"
	"music/log"
	"testing"
)

func TestMusicModule_SearchMusic(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewCloudAPI(Loggers, "")
	res, err := Api.Music().SearchMusic(types.SearchParams{
		Keywords: "大海",
		Type:     1,
		Limit:    30,
		Offset:   0,
	})
	if err != nil {
		t.Error(err)
	}
	type Song struct {
		Result struct {
			Songs []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"songs"`
			SongCount int `json:"songCount"`
		} `json:"result"`
		Code int `json:"code"`
	}
	var b Song
	fmt.Println(res)
	err = json.Unmarshal(res, &b)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	t.Log(b)
}
