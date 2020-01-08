package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/euronetzrt/xenstat-exporter/xenstat"
	"strconv"
)

type networkCollector struct {
	rbytes   *prometheus.Desc
	rpackets *prometheus.Desc
	rerrs    *prometheus.Desc
	rdrop    *prometheus.Desc
	tbytes   *prometheus.Desc
	tpackets *prometheus.Desc
	terrs    *prometheus.Desc
	tdrop    *prometheus.Desc
}

func newNetworkCollector() networkCollector {
	labels := []string{"domain", "network", "network_id"}

	return networkCollector{
		rbytes: prometheus.NewDesc(
			"xenstat_network_rbytes",
			"number of receive bytes for this network",
			labels,
			nil,
		),
		rpackets: prometheus.NewDesc(
			"xenstat_network_rpackets",
			"number of receive packets for this network",
			labels,
			nil,
		),
		rerrs: prometheus.NewDesc(
			"xenstat_network_rerrs",
			"number of receive errors for this network",
			labels,
			nil,
		),
		rdrop: prometheus.NewDesc(
			"xenstat_network_rdrop",
			"number of receive drops for this network",
			labels,
			nil,
		),
		tbytes: prometheus.NewDesc(
			"xenstat_network_tbytes",
			"number of transmit bytes for this network",
			labels,
			nil,
		),
		tpackets: prometheus.NewDesc(
			"xenstat_network_tpackets",
			"number of transmit packets for this network",
			labels,
			nil,
		),
		terrs: prometheus.NewDesc(
			"xenstat_network_terrs",
			"number of transmit errors for this network",
			labels,
			nil,
		),
		tdrop: prometheus.NewDesc(
			"xenstat_network_tdrop",
			"number of transmit drops for this network",
			labels,
			nil,
		),
	}
}

func (c networkCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.rbytes
	ch <- c.rpackets
	ch <- c.rerrs
	ch <- c.rdrop
	ch <- c.tbytes
	ch <- c.tpackets
	ch <- c.terrs
	ch <- c.tdrop
}

func (c networkCollector) Collect(domain string, n *xenstat.Network, ch chan<- prometheus.Metric) {
	network := strconv.FormatUint(uint64(n.Idx), 10)
	networkID := strconv.FormatUint(uint64(n.ID()), 10)

	ch <- prometheus.MustNewConstMetric(
		c.rbytes,
		prometheus.CounterValue,
		float64(n.RBytes()),
		domain,
		network,
		networkID,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rpackets,
		prometheus.CounterValue,
		float64(n.RPackets()),
		domain,
		network,
		networkID,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rerrs,
		prometheus.CounterValue,
		float64(n.RErrs()),
		domain,
		network,
		networkID,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rdrop,
		prometheus.CounterValue,
		float64(n.RDrop()),
		domain,
		network,
		networkID,
	)

	ch <- prometheus.MustNewConstMetric(
		c.tbytes,
		prometheus.CounterValue,
		float64(n.TBytes()),
		domain,
		network,
		networkID,
	)

	ch <- prometheus.MustNewConstMetric(
		c.tpackets,
		prometheus.CounterValue,
		float64(n.TPackets()),
		domain,
		network,
		networkID,
	)

	ch <- prometheus.MustNewConstMetric(
		c.terrs,
		prometheus.CounterValue,
		float64(n.TErrs()),
		domain,
		network,
		networkID,
	)

	ch <- prometheus.MustNewConstMetric(
		c.tdrop,
		prometheus.CounterValue,
		float64(n.TDrop()),
		domain,
		network,
		networkID,
	)
}
