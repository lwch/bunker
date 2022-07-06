package app

import (
	"github.com/kardianos/service"
	"github.com/lwch/bunker/code/conf"
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
	logging.SetSizeRotate(logging.SizeRotateConfig{
		Dir:         app.cfg.LogDir,
		Name:        "bunker",
		Size:        int64(app.cfg.LogSize.Bytes()),
		Rotate:      app.cfg.LogRotate,
		WriteStdout: true,
		WriteFile:   true,
	})
	defer logging.Flush()

	if len(app.cfg.Server) > 0 {
		app.runServer()
	} else {
		app.runAgent()
	}
}

// Stop stop application
func (a *App) Stop(s service.Service) error {
	return nil
}
