package serviceoverview

import (
	"net/http"
	"time"

	"github.com/CloudDetail/apo/backend/pkg/model/request"
	"github.com/CloudDetail/apo/backend/pkg/services/serviceoverview"

	"github.com/CloudDetail/apo/backend/pkg/code"
	"github.com/CloudDetail/apo/backend/pkg/core"

	"github.com/CloudDetail/apo/backend/pkg/model/response"
)

// GetEndPointsData 获取endpoints服务列表
// @Summary 获取endpoints服务列表
// @Description 获取endpoints服务列表
// @Tags API.service
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param startTime query int64 true "查询开始时间"
// @Param endTime query int64 true "查询结束时间"
// @Param step query int64 true "步长"
// @Param serviceName query string false "名称"
// @Param sortRule query int true "排序逻辑"
// @Success 200 {object} response.ServiceEndPointsRes
// @Failure 400 {object} code.Failure
// @Router /api/service/endpoints [get]
func (h *handler) GetEndPointsData() core.HandlerFunc {
	return func(c core.Context) {
		req := new(request.GetEndPointsDataRequest)
		if err := c.ShouldBindQuery(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		var startTime time.Time
		var endTime time.Time
		req.StartTime = req.StartTime / 1000000 //接收的微秒级别的startTime和endTime需要先转成秒级别
		req.EndTime = req.EndTime / 1000000     //接收的微秒级别的startTime和endTime需要先转成秒级别
		startTime = time.Unix(req.StartTime, 0)
		endTime = time.Unix(req.EndTime, 0)
		step := time.Duration(req.Step * 1000)
		//step := time.Minute
		serviceName := req.ServiceName
		sortRule := serviceoverview.SortType(req.SortRule)
		var res []response.ServiceEndPointsRes
		data, err := h.serviceoverview.GetServicesEndPointData(startTime, endTime, step, serviceName, sortRule)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.GetTop3UrlListError,
				code.Text(code.GetTop3UrlListError)).WithError(err),
			)
			return
		}
		if data != nil {
			res = data
		} else {
			res = []response.ServiceEndPointsRes{} // 确保返回一个空数组
		}

		c.Payload(res)
	}
}