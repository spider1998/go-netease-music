package api

import (
	"music/log"
)

type CloudAPI struct {
	log.Logger
	auth  TestModule
	user  UserModule
	music MusicModule
}

func NewCloudAPI(logger log.Logger, gateway string) *CloudAPI {
	api := &CloudAPI{
		Logger: logger,
	}
	api.auth = TestModule{api}
	api.user = UserModule{api}
	api.music = MusicModule{api}
	return api
}

func (api CloudAPI) Test() TestModule {
	return api.auth
}

func (api CloudAPI) User() UserModule {
	return api.user
}

func (api CloudAPI) Music() MusicModule {
	return api.music
}
