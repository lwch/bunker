package conf

import (
	"os"
	"path/filepath"

	"github.com/lwch/bunker/code/utils"
	"github.com/lwch/runtime"
	"github.com/lwch/yaml"
)

type Configure struct {
	ID        string
	Server    string
	UseSSL    bool
	Listen    uint16
	TLSKey    string
	TLSCrt    string
	LogDir    string
	LogSize   utils.Bytes
	LogRotate int
}

// Load load configure file
func Load(dir string) *Configure {
	var cfg struct {
		ID     string `yaml:"id"`
		Server string `yaml:"server"`
		SSL    bool   `yaml:"ssl"`
		Listen uint16 `yaml:"listen"`
		Secret string `yaml:"secret"`
		Log    struct {
			Dir    string      `yaml:"dir"`
			Size   utils.Bytes `yaml:"size"`
			Rotate int         `yaml:"rotate"`
		} `yaml:"log"`
		TLS struct {
			Key string `yaml:"key"`
			Crt string `yaml:"crt"`
		} `yaml:"tls"`
	}
	runtime.Assert(yaml.Decode(dir, &cfg))
	if !filepath.IsAbs(cfg.Log.Dir) {
		dir, err := os.Executable()
		runtime.Assert(err)
		cfg.Log.Dir = filepath.Join(filepath.Dir(dir), cfg.Log.Dir)
	}
	return &Configure{
		ID:        cfg.ID,
		Server:    cfg.Server,
		UseSSL:    cfg.SSL,
		Listen:    cfg.Listen,
		TLSKey:    cfg.TLS.Key,
		TLSCrt:    cfg.Secret,
		LogDir:    cfg.Log.Dir,
		LogSize:   cfg.Log.Size,
		LogRotate: cfg.Log.Rotate,
	}
}
