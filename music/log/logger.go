package log

import (
	"github.com/inconshreveable/log15"
)

type Logger interface {
	log15.Logger
}

func New(debug bool, module string) (Logger, error) {
	lvl := log15.LvlInfo
	if debug {
		lvl = log15.LvlDebug
	}
	logger := log15.Root()
	h := log15.LvlFilterHandler(lvl, log15.CallerFileHandler(logger.GetHandler()))
	logger.SetHandler(h)
	return logger, nil
}

type Message string

func (e Message) Error() string {
	return string(e)
}
