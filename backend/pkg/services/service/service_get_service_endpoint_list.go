package service

import "github.com/CloudDetail/apo/backend/pkg/model/request"

func (s *service) GetServiceEndPointList(req *request.GetServiceEndPointListRequest) ([]string, error) {
	// 获取服务Endpoint列表
	return s.promRepo.GetServiceEndPointList(req.StartTime, req.EndTime, req.ServiceName)
}