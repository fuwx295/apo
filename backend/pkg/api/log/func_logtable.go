package log

import (
	"net/http"

	"github.com/CloudDetail/apo/backend/pkg/code"
	"github.com/CloudDetail/apo/backend/pkg/core"
	"github.com/CloudDetail/apo/backend/pkg/model/request"
)

func (h *handler) CreateLogTable() core.HandlerFunc {
	return func(c core.Context) {
		req := new(request.LogTableRequest)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		req.FillerValue()
		resp, err := h.logService.CreateLogTable(req)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CreateLogTableError,
				code.Text(code.CreateLogTableError)).WithError(err),
			)
			return
		}
		c.Payload(resp)
	}
}

func (h *handler) DropLogTable() core.HandlerFunc {
	return func(c core.Context) {
		req := new(request.LogTableRequest)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		resp, err := h.logService.DropLogTable(req)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.DropLogTableError,
				code.Text(code.DropLogTableError)).WithError(err),
			)
			return
		}
		c.Payload(resp)
	}
}

func (h *handler) UpdateLogTable() core.HandlerFunc {
	return func(c core.Context) {
		req := new(request.LogTableRequest)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		resp, err := h.logService.UpdateLogTable(req)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UpateLogTableError,
				code.Text(code.UpateLogTableError)).WithError(err),
			)
			return
		}
		c.Payload(resp)
	}
}
