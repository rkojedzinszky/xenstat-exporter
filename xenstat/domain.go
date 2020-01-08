package xenstat

/*
#include <xenstat.h>
*/
import "C"

// Domain mirrors xenstat_domain
type Domain struct {
	d *C.xenstat_domain
}

// Name returns the Domain's name
func (d *Domain) Name() string {
	return C.GoString(C.xenstat_domain_name(d.d))
}

// CPUNs returns information about how much CPU time has been used
func (d *Domain)CPUNs() uint64 {
	return uint64(C.xenstat_domain_cpu_ns(d.d))
}

// NumVCPUs returns the number of VCPUs allocated to a domain
func (d *Domain) NumVCPUs() uint {
	return uint(C.xenstat_domain_num_vcpus(d.d))
}

// VCPU returns the VCPU handle to obtain VCPU stats
func (d *Domain)VCPU(idx uint) *VCPU {
	if v := C.xenstat_domain_vcpu(d.d, C.uint(idx)); v != nil {
		return &VCPU{
			Idx: idx,
			v:   v,
		}
	}

	return nil
}

// CurMem returns the current memory reservation for this domain
func (d *Domain)CurMem() uint64 {
	return uint64(C.xenstat_domain_cur_mem(d.d))
}

// MaxMem returns the maximum memory reservation for this domain
func (d *Domain)MaxMem() uint64 {
	return uint64(C.xenstat_domain_max_mem(d.d))
}

// SSID returns the domain's SSID
func (d *Domain)SSID() uint {
	return uint(C.xenstat_domain_ssid(d.d))
}

// NumNetworks returns the number of networks for a given domain
func (d *Domain)NumNetworks() uint {
	return uint(C.xenstat_domain_num_networks(d.d))
}

// Network returns the network handle to obtain network stats
func (d *Domain)Network(idx uint) *Network {
	if n := C.xenstat_domain_network(d.d, C.uint(idx)); n != nil {
		return &Network{
			Idx: idx,
			n:   n,
		}
	}

	return nil
}

// NumVBDs returns the number of VBDs for a given domain
func (d *Domain)NumVBDs() uint {
	return uint(C.xenstat_domain_num_vbds(d.d))
}

// VBD returns the VBD handle to obtain VBD stats
func (d *Domain)VBD(idx uint) *VBD {
	if v := C.xenstat_domain_vbd(d.d, C.uint(idx)); v != nil {
		return &VBD{
			Idx: idx,
			v:   v,
		}
	}
	
	return nil
}