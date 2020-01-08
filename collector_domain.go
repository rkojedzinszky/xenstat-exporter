package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/euronetzrt/xenstat-exporter/xenstat"
)

type domainCollector struct {
	cpuNs       *prometheus.Desc
	numVcpus    *prometheus.Desc
	curMem      *prometheus.Desc
	maxMem      *prometheus.Desc
	ssid        *prometheus.Desc
	numNetworks *prometheus.Desc
	numVBDs     *prometheus.Desc

	vcpu    vcpuCollector
	network networkCollector
	vbd     vbdCollector
}

func newDomainCollector() domainCollector {
	labels := []string{"domain"}

	return domainCollector{
		cpuNs: prometheus.NewDesc(
			"xenstat_domain_cpu_ns",
			"information about how much CPU time has been used",
			labels,
			nil,
		),
		numVcpus: prometheus.NewDesc(
			"xenstat_domain_num_vcpus",
			"number of VCPUs allocated to a domain",
			labels,
			nil,
		),
		curMem: prometheus.NewDesc(
			"xenstat_domain_cur_mem",
			"Current memory reservation for this domain",
			labels,
			nil,
		),
		maxMem: prometheus.NewDesc(
			"xenstat_domain_max_mem",
			"Maximum memory reservation for this domain",
			labels,
			nil,
		),
		ssid: prometheus.NewDesc(
			"xenstat_domain_ssid",
			"Domain's SSID",
			labels,
			nil,
		),
		numNetworks: prometheus.NewDesc(
			"xenstat_domain_num_networks",
			"number of networks for a given domain",
			labels,
			nil,
		),
		numVBDs: prometheus.NewDesc(
			"xenstat_domain_num_vbds",
			"number of VBDs for a given domain",
			labels,
			nil,
		),
		vcpu:    newVcpuCollector(),
		network: newNetworkCollector(),
		vbd:     newVBDCollector(),
	}
}

func (c domainCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.cpuNs
	ch <- c.numVcpus
	ch <- c.curMem
	ch <- c.maxMem
	ch <- c.ssid
	ch <- c.numNetworks
	ch <- c.numVBDs

	c.vcpu.Describe(ch)
	c.network.Describe(ch)
	c.vbd.Describe(ch)
}

func (c domainCollector) Collect(d *xenstat.Domain, ch chan<- prometheus.Metric) {
	name := d.Name()

	ch <- prometheus.MustNewConstMetric(
		c.cpuNs,
		prometheus.CounterValue,
		float64(d.CPUNs()),
		name,
	)

	numVcpus := d.NumVCPUs()
	ch <- prometheus.MustNewConstMetric(
		c.numVcpus,
		prometheus.GaugeValue,
		float64(numVcpus),
		name,
	)
	for i := uint(0); i < numVcpus; i++ {
		c.vcpu.Collect(name, d.VCPU(i), ch)
	}

	ch <- prometheus.MustNewConstMetric(
		c.curMem,
		prometheus.GaugeValue,
		float64(d.CurMem()),
		name,
	)

	ch <- prometheus.MustNewConstMetric(
		c.maxMem,
		prometheus.GaugeValue,
		float64(d.MaxMem()),
		name,
	)

	ch <- prometheus.MustNewConstMetric(
		c.ssid,
		prometheus.GaugeValue,
		float64(d.SSID()),
		name,
	)

	numNetworks := d.NumNetworks()
	ch <- prometheus.MustNewConstMetric(
		c.numNetworks,
		prometheus.GaugeValue,
		float64(numNetworks),
		name,
	)
	for i := uint(0); i < numNetworks; i++ {
		c.network.Collect(name, d.Network(i), ch)
	}

	numVBDs := d.NumVBDs()
	ch <- prometheus.MustNewConstMetric(
		c.numVBDs,
		prometheus.GaugeValue,
		float64(numVBDs),
		name,
	)
	for i := uint(0); i < numVBDs; i++ {
		c.vbd.Collect(name, d.VBD(i), ch)
	}
}
