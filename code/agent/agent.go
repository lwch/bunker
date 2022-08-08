package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/lwch/bunker/code/agent/shell"
	"github.com/lwch/bunker/code/agent/vnc"
	"github.com/lwch/bunker/code/conf"
	"github.com/lwch/bunker/code/network"
	"github.com/lwch/bunker/code/utils"
	"github.com/lwch/logging"
	"github.com/lwch/runtime"
	"github.com/shirou/gopsutil/v3/host"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type agent struct {
	cfg     *conf.Configure
	id      string
	workDir string
	version string // TODO
}

func newAgent(cfg *conf.Configure, workDir string) *agent {
	return &agent{
		cfg:     cfg,
		workDir: workDir,
	}
}

// Run run in agent mode
func (a *agent) Run() {
	var options []grpc.DialOption
	if !a.cfg.UseSSL {
		options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	options = append(options, grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = a.cfg.BuildHeader(ctx, a.id)
		return invoker(ctx, method, req, reply, cc, opts...)
	}))

	conn, err := grpc.Dial(a.cfg.Server, options...)
	runtime.Assert(err)
	defer conn.Close()

	logging.Info("connected on %s", a.cfg.Server)

	cli := network.NewBunkerClient(conn)

	logging.Info("waiting for handshake...")
	err = a.handshake(cli)
	if err != nil {
		logging.Error("wait for handshake: %v", err)
		return
	}
	logging.Info("handshake ok, agent_id=%s", a.id)

	ctx, cancel := context.WithCancel(context.Background())

	loop := func(name string, cb func(context.Context) error) {
		defer utils.Recover(name)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
			default:
				err := cb(ctx)
				if err != nil {
					logging.Error("call %s failed: %s", err)
					return
				}
			}
		}
	}

	sh := shell.New(a.cfg, cli)
	vn := vnc.New(a.cfg, cli)

	go loop("run_shell", sh.Run)
	go loop("shell_resize", sh.Resize)
	go loop("run_vnc", vn.Run)

	<-ctx.Done()
	os.Exit(1)
}

func (a *agent) handshake(cli network.BunkerClient) error {
	id := a.getID()
	info, _ := host.Info()
	args := network.ConnectArgs{
		Id:       id,
		Version:  a.version,
		Hostname: info.Hostname,
		Os:       info.OS,
		Platform: info.Platform,
		Arch:     info.KernelArch,
	}
	resp, err := cli.Connect(context.Background(), &args)
	if err != nil {
		return err
	}
	if id != resp.GetId() && len(resp.GetId()) > 0 {
		err = a.setID(resp.GetId())
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *agent) getID() string {
	dir := filepath.Join(a.workDir, "id")
	data, _ := os.ReadFile(dir)
	a.id = strings.TrimSpace(string(data))
	return a.id
}

func (a *agent) setID(id string) error {
	dir := filepath.Join(a.workDir, "id")
	a.id = id
	return os.WriteFile(dir, []byte(id), 0644)
}
