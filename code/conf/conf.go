package conf

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/lwch/bunker/code/utils"
	"github.com/lwch/runtime"
	"github.com/lwch/yaml"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Configure struct {
	Server     string
	UseSSL     bool
	GrpcListen uint16
	HttpListen uint16
	TLSKey     string
	TLSCrt     string
	LogDir     string
	LogSize    utils.Bytes
	LogRotate  int
	otp        *utils.TOTP
}

// Load load configure file
func Load(dir string) *Configure {
	var cfg struct {
		Server     string `yaml:"server"`
		SSL        bool   `yaml:"ssl"`
		GrpcListen uint16 `yaml:"grpc_listen"`
		HttpListen uint16 `yaml:"http_listen"`
		Secret     string `yaml:"secret"`
		Log        struct {
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
		Server:     cfg.Server,
		UseSSL:     cfg.SSL,
		GrpcListen: cfg.GrpcListen,
		HttpListen: cfg.HttpListen,
		TLSKey:     cfg.TLS.Key,
		TLSCrt:     cfg.Secret,
		LogDir:     cfg.Log.Dir,
		LogSize:    cfg.Log.Size,
		LogRotate:  cfg.Log.Rotate,
		otp:        utils.NewTOTP(cfg.Secret),
	}
}

// BuildHeader build request header
func (cfg *Configure) BuildHeader(ctx context.Context, id string) context.Context {
	if len(id) > 0 {
		ctx = metadata.AppendToOutgoingContext(ctx, "id", id)
	}
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "TOTP "+cfg.otp.Gen())
	return ctx
}

// SecretVerify
func (cfg *Configure) SecretVerify(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}
	token := md.Get("authorization")
	if len(token) == 0 {
		return status.Errorf(codes.Unauthenticated, "Missing authorization token")
	}
	tk := token[0]
	if !strings.HasPrefix(tk, "TOTP ") {
		return status.Error(codes.Unauthenticated, "Invalid token")
	}
	tk = strings.TrimPrefix(tk, "TOTP ")
	if !cfg.otp.Verify(tk) {
		return status.Error(codes.Unauthenticated, "Invalid token")
	}
	return nil
}
