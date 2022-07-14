package app

import (
	"os"
	rt "runtime"

	"github.com/kardianos/service"
	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/logging"
)

type handler interface {
	Run()
}

// App main instance
type App struct {
	h    handler
	cfg  *conf.Configure
	name string
}

func New(h handler, cfg *conf.Configure, name string) *App {
	return &App{h: h, cfg: cfg, name: name}
}

// Start start application
func (app *App) Start(s service.Service) error {
	go app.run()
	return nil
}

func (app *App) run() {
	stdout := true
	if rt.GOOS == "windows" {
		stdout = false
	}
	logging.SetSizeRotate(logging.SizeRotateConfig{
		Dir:         app.cfg.LogDir,
		Name:        app.name,
		Size:        int64(app.cfg.LogSize.Bytes()),
		Rotate:      app.cfg.LogRotate,
		WriteStdout: stdout,
		WriteFile:   true,
	})
	defer logging.Flush()

	app.h.Run()

	logging.Flush()
	os.Exit(1)
}

// Stop stop application
func (a *App) Stop(s service.Service) error {
	return nil
}
