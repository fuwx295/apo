package service

import (
	"go.uber.org/zap"

	"github.com/CloudDetail/apo/backend/pkg/core"
	"github.com/CloudDetail/apo/backend/pkg/repository/clickhouse"
	"github.com/CloudDetail/apo/backend/pkg/repository/database"
	"github.com/CloudDetail/apo/backend/pkg/repository/polarisanalyzer"
	"github.com/CloudDetail/apo/backend/pkg/repository/prometheus"
	"github.com/CloudDetail/apo/backend/pkg/services/service"
)

type Handler interface {
	// GetServiceEndpointTopology 获取服务上下游拓扑
	// @Tags API.service
	// @Router /api/service/topology [get]
	GetServiceEndpointTopology() core.HandlerFunc

	// GetDescendantMetrics 获取所有下游服务的延时曲线数据
	// @Tags API.service
	// @Router /api/service/descendant/metrics [get]
	GetDescendantMetrics() core.HandlerFunc

	// GetDescendantRelevance 获取依赖节点延时关联度
	// @Tags API.service
	// @Router /api/service/descendant/relevance [get]
	GetDescendantRelevance() core.HandlerFunc

	// GetPolarisInfer 获取北极星指标分析情况
	// @Tags API.service
	// @Router /api/service/polaris/infer [get]
	GetPolarisInfer() core.HandlerFunc

	// GetErrorInstance 获取错误实例
	// @Tags API.service
	// @Router /api/service/error/instance [get]
	GetErrorInstance() core.HandlerFunc

	// GetErrorInstanceLogs 获取错误实例故障现场日志
	// @Tags API.service
	// @Router /api/service/errorinstance/logs [get]
	GetErrorInstanceLogs() core.HandlerFunc

	// GetLogMetrics 获取日志相关指标
	// @Tags API.service
	// @Router /api/service/log/metrics [get]
	GetLogMetrics() core.HandlerFunc

	// GetLogLogs 获取Log故障现场日志
	// @Tags API.service
	// @Router /api/service/log/logs [get]
	GetLogLogs() core.HandlerFunc

	// GetTraceMetrics 获取Trace相关指标
	// @Tags API.service
	// @Router /api/service/trace/metrics [get]
	GetTraceMetrics() core.HandlerFunc

	// GetTraceLogs 获取Trace故障现场日志
	// @Tags API.service
	// @Router /api/service/trace/logs [get]
	GetTraceLogs() core.HandlerFunc

	// GetServiceList 获取服务列表
	// @Tags API.service
	// @Router /api/service/list [get]
	GetServiceList() core.HandlerFunc

	// GetServiceInstanceList 获取服务实例列表
	// @Tags API.service
	// @Router /api/service/instance/list [get]
	GetServiceInstanceList() core.HandlerFunc

	// GetServiceInstanceOptions 获取服务下拉实例列表
	// @Tags API.service
	// @Router /api/service/instance/options [get]
	GetServiceInstanceOptions() core.HandlerFunc

	// GetServiceEndPointList 获取服务EndPoint列表
	// @Tags API.service
	// @Router /api/service/endpoint/list [get]
	GetServiceEndPointList() core.HandlerFunc

	// CountK8sEvents 获取K8s事件数量
	// @Tags API.service
	// @Router /api/service/k8s/events/count [get]
	CountK8sEvents() core.HandlerFunc
}

type handler struct {
	logger             *zap.Logger
	serviceInfoService service.Service
}

func New(logger *zap.Logger, chRepo clickhouse.Repo, promRepo prometheus.Repo, polRepo polarisanalyzer.Repo, dbRepo database.Repo) Handler {
	return &handler{
		logger:             logger,
		serviceInfoService: service.New(chRepo, promRepo, polRepo, dbRepo),
	}
}