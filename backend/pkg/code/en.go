package code

var enText = map[string]string{
	ServerError:    "Internal server error",
	ParamBindError: "Parameter error",
	DbConnectError: "Failed to connect Database",

	MockCreateError: "Failed to create mock",
	MockListError:   "Failed to get mock list",
	MockDetailError: "Failed to get mock detail",
	MockDeleteError: "Failed to delete mock",

	GetServiceUrlTopologyError:     "Failed to get service url topology",
	GetDescendantMetricsError:      "Failed to get descendant metrics",
	GetDescendantRelevanceError:    "Failed to get descendant relevance",
	GetPolarisInferError:           "Failed to get polaris infer",
	GetErrorInstanceError:          "Failed to get error instance",
	GetErrorInstanceLogsError:      "Failed to get error instance logs",
	GetLogMetricsError:             "Failed to get log metrics",
	GetLogLogsError:                "Failed to get log logs",
	GetTraceMetricsError:           "Failed to get trace metrics",
	GetTraceLogsError:              "Failed to get trace logs",
	GetServiceListError:            "Failed to get service list",
	GetServiceInstanceOptionsError: "Failed to get service instance list",
	GetK8sEventError:               "Failed to get k8s events",
	GetServiceEndPointListError:    "Failed to get service endpoint list",

	GetFaultLogPageListError: "Failed to get faultlog pagelist",
	GetFaultLogContentError:  "Failed to get faultlog content",

	GetTracePageListError: "Failed to get trace pagelist",

	GetOverviewServiceInstanceListError: "Failed to get overview service instance list",
	GetServiceMoreUrlListError:          "Failed to get service more url list",
	GetThresholdError:                   "Failed to get threshold",
	GetTop3UrlListError:                 "Failed to get top3 url list",
	SetThresholdError:                   "Failed to set threshold",
	GetServicesAlertError:               "Failed to get services alert",
}