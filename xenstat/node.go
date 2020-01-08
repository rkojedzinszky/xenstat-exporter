package xenstat

/*
#include <xenstat.h>
*/
import "C"

// Node mirrors xenstat_node
type Node struct {
	n *C.xenstat_node
}

// GetNode returns a Node struct
func (h *Handle) GetNode() *Node {
	if n := C.xenstat_get_node(h.h, C.XENSTAT_ALL); n != nil {
		return &Node{n: n}
	}

	return nil
}

// Free frees a Node struct
func (n *Node) Free() {
	C.xenstat_free_node(n.n)
}

// XenVersion returns xen version of the node
func (n *Node) XenVersion() string {
	return C.GoString(C.xenstat_node_xen_version(n.n))
}

// TotMem returns the amount of total memory on a node
func (n *Node)TotMem() uint64 {
	return uint64(C.xenstat_node_tot_mem(n.n))
}

// FreeMem returns the amount of free memory on a node
func (n *Node)FreeMem() uint64 {
	return uint64(C.xenstat_node_free_mem(n.n))
}

// NumDomains returns the number of domains existing on a node
func (n *Node) NumDomains() uint {
	return uint(C.xenstat_node_num_domains(n.n))
}

// DomainByIndex returns the Domain indexed by idx
func (n *Node) DomainByIndex(idx uint) *Domain {
	if d := C.xenstat_node_domain_by_index(n.n, C.uint(idx)); d != nil {
		return &Domain{d: d}
	}

	return nil
}

// NumCpus returns the number of CPUs existing on a node
func (n *Node) NumCpus() uint {
	return uint(C.xenstat_node_num_cpus(n.n));
}

// CPUHz returns information about the CPU speed
func (n *Node) CPUHz() uint64 {
	return uint64(C.xenstat_node_cpu_hz(n.n))
}
