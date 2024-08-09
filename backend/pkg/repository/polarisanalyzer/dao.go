package polarisanalyzer

import (
	"net/http"
	"os"
	"time"
)

var polarisAnalyzerAddress = "http://localhost:5000"

type Repo interface {
	// SortDescendantByLatencyRelevance 查询依赖节点延时关联度
	SortDescendantByLatencyRelevance(
		startTime, endTime int64, stepStr string,
		targetService, targetEndpoint string,
		unsortedDescendant []LatencyRelevance,
	) (sorted []LatencyRelevance, unsorted []LatencyRelevance, err error)

	QueryPolarisInfer(
		startTime, endTime int64, stepStr string,
		service, endpoint string,
	) (*PolarisInferRes, error)
}

func New() (Repo, error) {
	if value, find := os.LookupEnv("POLARIS_ANALYZER_ADDRESS"); find {
		polarisAnalyzerAddress = value
	}

	return &polRepo{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

type polRepo struct {
	client *http.Client
}