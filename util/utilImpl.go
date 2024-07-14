package util

import (
	"github.com/chaaaeeee/sireng/config"
)

type utilImpl struct {
	config *config.Config
}

func NewUtil(config *config.Config) Util {
	return &utilImpl{config: config}
}
