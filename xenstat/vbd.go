package xenstat

/*
#include <xenstat.h>
*/
import "C"

type VBD struct {
	Idx uint
	v *C.xenstat_vbd
}

// Type returns the back driver type  for Virtual Block Device
func (v *VBD)Type() uint {
	return uint(C.xenstat_vbd_type(v.v))
}

// Dev returns the device number for Virtual Block Device
func (v *VBD)Dev() uint {
	return uint(C.xenstat_vbd_dev(v.v))
}

// OOReqs returns the number of OO requests for vbd
func (v *VBD)OOReqs() uint64 {
	return uint64(C.xenstat_vbd_oo_reqs(v.v))
}

// RDReqs returns the number of RD requests for vbd
func (v *VBD)RDReqs() uint64 {
	return uint64(C.xenstat_vbd_rd_reqs(v.v))
}

// WRReqs returns the number of WR requests for vbd
func (v *VBD)WRReqs() uint64 {
	return uint64(C.xenstat_vbd_wr_reqs(v.v))
}

// RDSects returns the number of RD Sectors for vbd
func (v *VBD)RDSects() uint64 {
	return uint64(C.xenstat_vbd_rd_sects(v.v))
}

// WRSects returns the number of WR Sectors for vbd
func (v *VBD)WRSects() uint64 {
	return uint64(C.xenstat_vbd_wr_sects(v.v))
}
