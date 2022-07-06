package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	rt "runtime"

	"github.com/kardianos/service"
	"github.com/lwch/bunker/code/app"
	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/runtime"
)

var (
	version      string = "0.0.0"
	gitHash      string
	gitReversion string
	buildTime    string
)

func showVersion() {
	fmt.Printf("version: v%s\ntime: %s\ncommit: %s.%s\n",
		version,
		buildTime,
		gitHash, gitReversion)
	os.Exit(0)
}

func main() {
	user := flag.String("user", "", "daemon user")
	cf := flag.String("conf", "", "configure file path")
	version := flag.Bool("version", false, "show version info")
	act := flag.String("action", "", "install or uninstall")
	flag.Parse()

	if *version {
		showVersion()
		os.Exit(0)
	}

	if len(*cf) == 0 {
		fmt.Println("missing -conf param")
		os.Exit(1)
	}

	dir, err := filepath.Abs(*cf)
	runtime.Assert(err)

	var depends []string
	if rt.GOOS != "windows" {
		depends = append(depends, "After=network.target")
	}

	appCfg := &service.Config{
		Name:         "np-svr",
		DisplayName:  "np-svr",
		Description:  "nat forward service",
		UserName:     *user,
		Arguments:    []string{"-conf", dir},
		Dependencies: depends,
	}

	cfg := conf.Load(*cf)

	app := app.New(cfg)
	sv, err := service.New(app, appCfg)
	runtime.Assert(err)

	switch *act {
	case "install":
		runtime.Assert(sv.Install())
	case "uninstall":
		runtime.Assert(sv.Uninstall())
	default:
		runtime.Assert(sv.Run())
	}
}
