package app

import (
	rt "runtime"

	"github.com/kardianos/service"
	"github.com/lwch/bunker/code/agent"
	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/bunker/code/server"
	"github.com/lwch/logging"
)

// App main instance
type App struct {
	cfg *conf.Configure
}

func New(cfg *conf.Configure) *App {
	return &App{cfg}
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
		Name:        "bunker",
		Size:        int64(app.cfg.LogSize.Bytes()),
		Rotate:      app.cfg.LogRotate,
		WriteStdout: stdout,
		WriteFile:   true,
	})
	defer logging.Flush()

	if len(app.cfg.Server) > 0 {
		agent.Run(app.cfg)
	} else {
		server.Run(app.cfg)
	}
}

// Stop stop application
func (a *App) Stop(s service.Service) error {
	return nil
}
