package shell

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lwch/bunker/code/network"
	"github.com/lwch/bunker/code/server/api"
	"github.com/lwch/logging"
)

func (h *Handler) resize(c *gin.Context) {
	logging.Info("resize")
	id := c.Param("id")
	cid := c.Param("cid")
	h.RLock()
	ch := h.waitResize[id]
	h.RUnlock()
	if ch == nil {
		api.NotFound(c, "agent")
		return
	}
	rows := c.GetUint("rows")
	cols := c.GetUint("cols")
	args := &network.ShellResizeArguments{
		Id:   cid,
		Rows: uint32(rows),
		Cols: uint32(cols),
	}
	select {
	case ch <- args:
		api.OK(c, nil)
	case <-time.After(api.SendTimeout):
		api.Timeout(c)
	}
}
