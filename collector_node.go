package main

import "fmt"
import "github.com/euronetzrt/xenstat-exporter/xenstat"
import "github.com/prometheus/client_golang/prometheus"

type nodeCollector struct {
	xenVersion *prometheus.Desc
	totMem     *prometheus.Desc
	freeMem    *prometheus.Desc
	numDomains *prometheus.Desc
	numCpus    *prometheus.Desc
	cpuHz      *prometheus.Desc

	domain domainCollector
}

func newNodeCollector() prometheus.Collector {
	return nodeCollector{
		xenVersion: prometheus.NewDesc(
			"xenstat_node_xen_version",
			"Running XEN version",
			[]string{"version"},
			nil,
		),
		totMem: prometheus.NewDesc(
			"xenstat_node_tot_mem",
			"Amount of total memory on a node",
			nil,
			nil,
		),
		freeMem: prometheus.NewDesc(
			"xenstat_node_free_mem",
			"Amount of free memory on a node",
			nil,
			nil,
		),
		numDomains: prometheus.NewDesc(
			"xenstat_node_num_domains",
			"Number of domains existing on a node",
			nil,
			nil,
		),
		numCpus: prometheus.NewDesc(
			"xenstat_node_num_cpus",
			"Number of CPUs existing on a node",
			nil,
			nil,
		),
		cpuHz: prometheus.NewDesc(
			"xenstat_node_cpu_hz",
			"Information about the CPU speed",
			nil,
			nil,
		),

		domain: newDomainCollector(),
	}
}

func (c nodeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.xenVersion
	ch <- c.totMem
	ch <- c.freeMem
	ch <- c.numDomains
	ch <- c.numCpus
	ch <- c.cpuHz

	c.domain.Describe(ch)
}

func (c nodeCollector) Collect(ch chan<- prometheus.Metric) {
	handle := xenstat.Init()
	if handle == nil {
		ch <- prometheus.NewInvalidMetric(c.xenVersion, fmt.Errorf("xenstat.Init() failed"))
		return
	}
	defer handle.Uninit()

	node := handle.GetNode()
	if node == nil {
		ch <- prometheus.NewInvalidMetric(c.xenVersion, fmt.Errorf("xenstat.GetNode() failed"))
		return
	}
	defer node.Free()

	ch <- prometheus.MustNewConstMetric(
		c.xenVersion,
		prometheus.GaugeValue,
		1,
		node.XenVersion(),
	)

	ch <- prometheus.MustNewConstMetric(
		c.totMem,
		prometheus.GaugeValue,
		float64(node.TotMem()),
	)

	ch <- prometheus.MustNewConstMetric(
		c.freeMem,
		prometheus.GaugeValue,
		float64(node.FreeMem()),
	)

	numDomains := node.NumDomains()
	ch <- prometheus.MustNewConstMetric(
		c.numDomains,
		prometheus.GaugeValue,
		float64(numDomains),
	)

	ch <- prometheus.MustNewConstMetric(
		c.numCpus,
		prometheus.GaugeValue,
		float64(node.NumCpus()),
	)

	ch <- prometheus.MustNewConstMetric(
		c.cpuHz,
		prometheus.GaugeValue,
		float64(node.CPUHz()),
	)

	for i := uint(0); i < numDomains; i++ {
		c.domain.Collect(node.DomainByIndex(i), ch)
	}
}
