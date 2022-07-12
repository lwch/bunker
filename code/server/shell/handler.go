package shell

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lwch/bunker/code/network"
)

// Handler shell handler
type Handler struct {
	sync.RWMutex
	waitRun    map[string]chan *network.RunShellArguments
	waitResize map[string]chan *network.ShellResizeArguments
}

// New create shell handler
func New() *Handler {
	return &Handler{
		waitRun:    make(map[string]chan *network.RunShellArguments),
		waitResize: make(map[string]chan *network.ShellResizeArguments),
	}
}

// ApiFuncs return api functions
func (h *Handler) ApiFuncs() []gin.RouteInfo {
	return []gin.RouteInfo{
		{Method: http.MethodConnect, Path: "/shell/:id", HandlerFunc: h.new},
		{Method: http.MethodPatch, Path: "/shell/:id/:cid/resize", HandlerFunc: h.resize},
	}
}

// WaitRun wait run by agent id
func (h *Handler) WaitRun(id string) <-chan *network.RunShellArguments {
	h.Lock()
	defer h.Unlock()
	if ch, ok := h.waitRun[id]; ok {
		return ch
	}
	h.waitRun[id] = make(chan *network.RunShellArguments)
	return h.waitRun[id]
}

// WaitResize wait resize by agent id
func (h *Handler) WaitResize(id string) <-chan *network.ShellResizeArguments {
	h.Lock()
	defer h.Unlock()
	if ch, ok := h.waitResize[id]; ok {
		return ch
	}
	h.waitResize[id] = make(chan *network.ShellResizeArguments)
	return h.waitResize[id]
}
