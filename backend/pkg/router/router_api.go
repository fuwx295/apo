package router

import (
	"github.com/CloudDetail/apo/backend/internal/api/mock"
	"github.com/CloudDetail/apo/backend/pkg/api/alerts"
	"github.com/CloudDetail/apo/backend/pkg/api/log"
	"github.com/CloudDetail/apo/backend/pkg/api/service"
	"github.com/CloudDetail/apo/backend/pkg/api/serviceoverview"
	"github.com/CloudDetail/apo/backend/pkg/api/trace"
)

func setApiRouter(r *resource) {
	api := r.mux.Group("/api")
	{
		mockHandler := mock.New(r.logger, r.internal_db)
		api.POST("/mock", mockHandler.Create())
		api.GET("/mock", mockHandler.List())
		api.GET("/mock/:id", mockHandler.Detail())
		api.DELETE("/mock/:id", mockHandler.Delete())
	}

	serviceApi := r.mux.Group("/api/service")
	{
		serviceOverviewHandler := serviceoverview.New(r.logger, r.ch, r.prom, r.pkg_db)
		serviceApi.GET("/endpoints", serviceOverviewHandler.GetEndPointsData())
		serviceApi.GET("/servicesAlert", serviceOverviewHandler.GetServicesAlert())
		serviceApi.GET("/moreUrl", serviceOverviewHandler.GetServiceMoreUrlList())
		serviceApi.GET("/instances", serviceOverviewHandler.GetServiceInstanceList())
		serviceApi.GET("/getThreshold", serviceOverviewHandler.GetThreshold())
		serviceApi.POST("/setThreshold", serviceOverviewHandler.SetThreshold())

		serviceHandler := service.New(r.logger, r.ch, r.prom, r.pol, r.pkg_db)
		serviceApi.GET("/topology", serviceHandler.GetServiceEndpointTopology())
		serviceApi.GET("/descendant/metrics", serviceHandler.GetDescendantMetrics())
		serviceApi.GET("/descendant/relevance", serviceHandler.GetDescendantRelevance())
		serviceApi.GET("/polaris/infer", serviceHandler.GetPolarisInfer())
		serviceApi.GET("/error/instance", serviceHandler.GetErrorInstance())
		serviceApi.GET("/errorinstance/logs", serviceHandler.GetErrorInstanceLogs())
		serviceApi.GET("/log/metrics", serviceHandler.GetLogMetrics())
		serviceApi.GET("/log/logs", serviceHandler.GetLogLogs())
		serviceApi.GET("/trace/metrics", serviceHandler.GetTraceMetrics())
		serviceApi.GET("/trace/logs", serviceHandler.GetTraceLogs())

		serviceApi.GET("/list", serviceHandler.GetServiceList())
		serviceApi.GET("/instance/list", serviceHandler.GetServiceInstanceList())
		serviceApi.GET("/instance/options", serviceHandler.GetServiceInstanceOptions())
		serviceApi.GET("/endpoint/list", serviceHandler.GetServiceEndPointList())
		serviceApi.GET("/k8s/events/count", serviceHandler.CountK8sEvents())
	}

	logApi := r.mux.Group("/api/log")
	{
		logHandler := log.New(r.logger, r.ch)
		logApi.POST("/fault/pagelist", logHandler.GetFaultLogPageList())
		logApi.POST("/fault/content", logHandler.GetFaultLogContent())
	}

	traceApi := r.mux.Group("/api/trace")
	{
		traceHandler := trace.New(r.logger, r.ch)
		traceApi.POST("/pagelist", traceHandler.GetTracePageList())
	}

	alertApi := r.mux.Group("/api/alerts")
	{
		alertHandler := alerts.New(r.logger, r.ch)
		alertApi.POST("/inputs/alertmanager", alertHandler.InputAlertManager())
	}
}