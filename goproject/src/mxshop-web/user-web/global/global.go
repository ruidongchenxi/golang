package global

import (
	"mxshop-web/user-web/config"

	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig *config.ServerConfig =&config.ServerConfig{}
    Trans ut.Translator
)