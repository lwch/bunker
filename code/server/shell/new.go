package shell

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lwch/bunker/code/network"
	"github.com/lwch/bunker/code/server/api"
	"github.com/lwch/runtime"
)

func (h *Handler) new(c *gin.Context) {
	id := c.Param("id")
	h.RLock()
	ch := h.waitRun[id]
	h.RUnlock()
	if ch == nil {
		api.NotFound(c, "agent")
		return
	}
	cid, err := runtime.UUID(16, "0123456789abcdef")
	runtime.Assert(err)
	args := &network.RunShellArguments{
		Id: cid,
	}
	select {
	case ch <- args:
		api.OK(c, cid)
	case <-time.After(api.SendTimeout):
		api.Timeout(c)
	}
}
