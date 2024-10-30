package service

import (
	"github.com/CloudDetail/apo/backend/pkg/model"
	"github.com/CloudDetail/apo/backend/pkg/model/response"
	"github.com/CloudDetail/apo/backend/pkg/repository/database"
	"github.com/CloudDetail/apo/backend/pkg/repository/prometheus"
	"github.com/CloudDetail/apo/backend/pkg/services/serviceoverview"
	"log"
	"math"
	"strconv"
	"time"
)

func (s *service) GetInstancesNew(startTime time.Time, endTime time.Time, step time.Duration, serviceName string, endPoint string) (res response.InstancesRes, err error) {
	threshold, err := s.dbRepo.GetOrCreateThreshold("", "", database.GLOBAL)
	if err != nil {
		return res, err
	}
	errorThreshold := threshold.ErrorRate
	tpsThreshold := threshold.Tps
	latencyThreshold := threshold.Latency

	filter := InstancesFilter{SrvName: serviceName, ContentKey: endPoint}
	filters := filter.ExtractFilterStr()

	// 获取RED指标
	instances := s.InstanceRED(startTime, endTime, filters)

	// 填充图表数据
	chartErr := s.InstanceRangeData(instances, startTime, endTime, step, filters)
	if chartErr.ErrorOrNil() != nil {
		log.Println("get instance range data error: ", chartErr)
	}
	// 填充日志数据
	logErr := s.InstanceLog(instances, startTime, endTime, step)
	if logErr.ErrorOrNil() != nil {
		log.Println("get instance log data error: ", logErr)
	}
	resData := []response.InstanceData{}
	res.Status = model.STATUS_NORMAL
	// 填充instance和RED指标的状态
	for _, instance := range instances.MetricGroupList {
		if instance.REDMetrics.DOD.ErrorRate != nil && *instance.REDMetrics.DOD.ErrorRate > errorThreshold {
			res.Status = model.STATUS_CRITICAL
		}
		if instance.REDMetrics.DOD.Latency != nil && *instance.REDMetrics.DOD.Latency > latencyThreshold {
			res.Status = model.STATUS_CRITICAL
		}
		if instance.REDMetrics.DOD.TPM != nil && *instance.REDMetrics.DOD.TPM > tpsThreshold {
			res.Status = model.STATUS_CRITICAL
		}

		if instance.REDMetrics.WOW.ErrorRate != nil && *instance.REDMetrics.WOW.ErrorRate > errorThreshold {
			res.Status = model.STATUS_CRITICAL
		}
		if instance.REDMetrics.WOW.Latency != nil && *instance.REDMetrics.WOW.Latency > latencyThreshold {
			res.Status = model.STATUS_CRITICAL
		}
		if instance.REDMetrics.WOW.TPM != nil && *instance.REDMetrics.WOW.TPM > tpsThreshold {
			res.Status = model.STATUS_CRITICAL
		}
	}

	for _, instance := range instances.MetricGroupList {
		if len(instance.Pod) == 0 && len(instance.PID) == 0 && len(instance.ContainerId) == 0 {
			continue
		}
		if instance.REDMetrics.Avg.Latency == nil {
			continue
		}

		// 构造latency的返回值
		latencyTempChartObject := response.TempChartObject{
			Ratio: response.Ratio{
				DayOverDay:  instance.REDMetrics.DOD.Latency,
				WeekOverDay: instance.REDMetrics.WOW.Latency,
			},
			Value: instance.REDMetrics.Avg.Latency,
		}
		if instance.LatencyData != nil {
			latencyTempChartObject.ChartData = DataToChart(instance.LatencyData)
		} else {
			latencyTempChartObject.ChartData = FillChart(startTime, endTime, step)
		}

		// 构造error的返回值
		errorTempChartObject := response.TempChartObject{
			Ratio: response.Ratio{
				DayOverDay:  instance.REDMetrics.DOD.ErrorRate,
				WeekOverDay: instance.REDMetrics.WOW.ErrorRate,
			},
			Value: instance.REDMetrics.Avg.ErrorRate,
		}
		if errorTempChartObject.Value == nil {
			zero := new(float64)
			errorTempChartObject.Value = zero
		}
		if instance.ErrorRateData != nil {
			errorTempChartObject.ChartData = DataToChart(instance.ErrorRateData)
		} else {
			errorTempChartObject.ChartData = FillChart(startTime, endTime, step)
		}

		// 构造tps返回值
		tpsTempChartObject := response.TempChartObject{
			Ratio: response.Ratio{
				DayOverDay:  instance.REDMetrics.DOD.TPM,
				WeekOverDay: instance.REDMetrics.WOW.TPM,
			},
			Value: instance.REDMetrics.Avg.TPM,
		}
		if instance.TPMData != nil {
			tpsTempChartObject.ChartData = DataToChart(instance.TPMData)
		} else {
			tpsTempChartObject.ChartData = FillChart(startTime, endTime, step)
		}

		// 构造log返回值
		logTempChartObject := response.TempChartObject{
			Ratio: response.Ratio{
				DayOverDay:  instance.LogDayOverDay,
				WeekOverDay: instance.LogWeekOverWeek,
			},
			Value: instance.LogAVGData,
		}
		filters := ExtractLogFilter(instance.InstanceKey)
		var normalLog []prometheus.MetricResult
		if len(filters) > 0 {
			normalLog = s.GetNormalLog(startTime, endTime, filters)
		}

		if logTempChartObject.Value == nil && normalLog != nil {
			zero := new(float64)
			logTempChartObject.Value = zero
		}
		if instance.LogData != nil {
			logTempChartObject.ChartData = DataToChart(instance.LogData)
		} else {
			logTempChartObject.ChartData = FillChart(startTime, endTime, step)
		}

		newInstance := response.InstanceData{
			Name:        instance.InstanceKey.GenInstanceName(),
			Namespace:   instance.Namespace,
			NodeName:    instance.NodeName,
			NodeIP:      instance.NodeIP,
			Timestamp:   nil,
			Latency:     latencyTempChartObject,
			Tps:         tpsTempChartObject,
			ErrorRate:   errorTempChartObject,
			Logs:        logTempChartObject,
			AlertStatus: model.NORMAL_ALERT_STATUS,
			AlertReason: model.AlertReason{},
		}

		pidI64, err := strconv.ParseInt(instance.PID, 10, 64)
		if err != nil {
			pidI64 = -1
		}

		instanceSingleList := []*model.ServiceInstance{
			{
				ServiceName: serviceName,
				ContainerId: instance.ContainerId,
				PodName:     instance.Pod,
				Namespace:   instance.Namespace,
				NodeName:    instance.NodeName,
				Pid:         pidI64,
			},
		}
		// 填充告警状态
		newInstance.AlertStatusCH = serviceoverview.GetAlertStatusCH(
			s.chRepo, &newInstance.AlertReason, nil,
			nil, serviceName, instanceSingleList,
			startTime, endTime,
		)

		// 填充末次启动时间
		startTSmap, _ := s.promRepo.QueryProcessStartTime(startTime, endTime, instanceSingleList)
		latestStartTime := getLatestStartTime(startTSmap) * 1e6
		if latestStartTime > 0 {
			newInstance.Timestamp = &latestStartTime
		}

		resData = append(resData, newInstance)
	}
	res.Data = resData

	for _, data := range res.Data {
		if data.InfrastructureStatus == model.STATUS_CRITICAL {
			res.Status = model.STATUS_CRITICAL
			break
		}
		if data.NetStatus == model.STATUS_CRITICAL {
			res.Status = model.STATUS_CRITICAL
		}
		if data.K8sStatus == model.STATUS_CRITICAL {
			res.Status = model.STATUS_CRITICAL
		}
	}
	return
}

// DataToChart 将图表数据转为map
// 可以复用
func DataToChart(data []prometheus.Points) map[int64]float64 {
	chart := make(map[int64]float64)
	for _, item := range data {
		timestamp := item.TimeStamp
		value := item.Value
		if !math.IsInf(value, 1) {
			chart[timestamp] = value
		}
	}
	return chart
}

func FillChart(startTime, endTime time.Time, step time.Duration) map[int64]float64 {
	values := make(map[int64]float64)
	for ts := startTime.UnixMicro(); ts <= endTime.UnixMicro(); ts += step.Microseconds() {
		values[ts] = 0
	}
	return values
}

func ExtractLogFilter(instance prometheus.InstanceKey) []string {
	var filters []string
	if len(instance.Pod) > 0 {
		filters = make([]string, 2)
		filters[0] = prometheus.LogMetricPodRegexPQLFilter
		filters[1] = instance.Pod
	} else if len(instance.PID) > 0 && len(instance.NodeName) > 0 {
		filters = make([]string, 4)
		filters[0] = prometheus.LogMetricPidRegexPQLFilter
		filters[1] = instance.PID
		filters[2] = prometheus.LogMetricNodeRegexPQLFilter
		filters[3] = instance.NodeName
	}
	return filters
}