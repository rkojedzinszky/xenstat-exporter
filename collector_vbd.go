package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/euronetzrt/xenstat-exporter/xenstat"
	"strconv"
)

type vbdCollector struct {
	ooreqs  *prometheus.Desc
	rdreqs  *prometheus.Desc
	wrreqs  *prometheus.Desc
	rdsects *prometheus.Desc
	wrsects *prometheus.Desc
}

func newVBDCollector() vbdCollector {
	labels := []string{"domain", "vbd"}

	return vbdCollector{
		ooreqs: prometheus.NewDesc(
			"xenstat_vbd_oo_reqs",
			"number of OO requests for vbd",
			labels,
			nil,
		),
		rdreqs: prometheus.NewDesc(
			"xenstat_vbd_rd_reqs",
			"number of RD requests for vbd",
			labels,
			nil,
		),
		wrreqs: prometheus.NewDesc(
			"xenstat_vbd_wr_reqs",
			"number of WR requests for vbd",
			labels,
			nil,
		),
		rdsects: prometheus.NewDesc(
			"xenstat_vbd_rd_sects",
			"number of RD Sectors for vbd",
			labels,
			nil,
		),
		wrsects: prometheus.NewDesc(
			"xenstat_vbd_wr_sects",
			"number of WR Sectors for vbd",
			labels,
			nil,
		),
	}
}

func (c vbdCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.ooreqs
	ch <- c.rdreqs
	ch <- c.wrreqs
	ch <- c.rdsects
	ch <- c.wrsects
}

func (c vbdCollector) Collect(domain string, v *xenstat.VBD, ch chan<- prometheus.Metric) {
	vbd := strconv.FormatUint(uint64(v.Idx), 10)

	ch <- prometheus.MustNewConstMetric(
		c.ooreqs,
		prometheus.CounterValue,
		float64(v.OOReqs()),
		domain,
		vbd,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rdreqs,
		prometheus.CounterValue,
		float64(v.RDReqs()),
		domain,
		vbd,
	)

	ch <- prometheus.MustNewConstMetric(
		c.wrreqs,
		prometheus.CounterValue,
		float64(v.WRReqs()),
		domain,
		vbd,
	)

	ch <- prometheus.MustNewConstMetric(
		c.rdsects,
		prometheus.CounterValue,
		float64(v.RDSects()),
		domain,
		vbd,
	)

	ch <- prometheus.MustNewConstMetric(
		c.wrsects,
		prometheus.CounterValue,
		float64(v.WRSects()),
		domain,
		vbd,
	)
}
