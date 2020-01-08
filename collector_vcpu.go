package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/euronetzrt/xenstat-exporter/xenstat"
	"strconv"
)

type vcpuCollector struct {
	online *prometheus.Desc
	usage  *prometheus.Desc
}

func newVcpuCollector() vcpuCollector {
	labels := []string{"domain", "vcpu"}

	return vcpuCollector{
		online: prometheus.NewDesc(
			"xenstat_vcpu_online",
			"Reflects online state of a vCPU",
			labels,
			nil,
		),
		usage: prometheus.NewDesc(
			"xenstat_vcpu_ns",
			"vCPU usage",
			labels,
			nil,
		),
	}
}

func (c vcpuCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.online
	ch <- c.usage
}

func (c vcpuCollector) Collect(domain string, v *xenstat.VCPU, ch chan<- prometheus.Metric) {
	vcpu := strconv.FormatUint(uint64(v.Idx), 10)

	ch <- prometheus.MustNewConstMetric(
		c.online,
		prometheus.GaugeValue,
		float64(v.Online()),
		domain,
		vcpu,
	)

	ch <- prometheus.MustNewConstMetric(
		c.usage,
		prometheus.CounterValue,
		float64(v.Usage()),
		domain,
		vcpu,
	)
}
