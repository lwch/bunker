package conf

import (
	"time"

	"github.com/lwch/bunker/code/utils"
)

type Configure struct {
	Server       string
	Listen       uint16
	TLSKey       string
	TLSCrt       string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	LogDir       string
	LogSize      utils.Bytes
	LogRotate    int
}

// Load load configure file
func Load(dir string) *Configure {
	return &Configure{}
}
