// Code generated by monitor-code-gen. DO NOT EDIT.

package tcp

import (
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "logstash-tcp"

var groupSet = map[string]bool{}

var metricSet = map[string]monitors.MetricInfo{}

var defaultMetrics = map[string]bool{}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "logstash-tcp",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         true,
}
