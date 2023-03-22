// Code generated by monitor-code-gen. DO NOT EDIT.

package statsd

import (
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "collectd/statsd"

var groupSet = map[string]bool{}

var metricSet = map[string]monitors.MetricInfo{}

var defaultMetrics = map[string]bool{}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "collectd/statsd",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         true,
}
