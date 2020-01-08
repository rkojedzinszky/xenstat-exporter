package xenstat

/*
#include <xenstat.h>
*/
import "C"

// Network mirrors xenstat_network
type Network struct {
	Idx uint
	n *C.xenstat_network
}

// ID for this network
func (n *Network)ID() uint {
	return uint(C.xenstat_network_id(n.n))
}

// RBytes returns the number of receive bytes for this network
func (n *Network)RBytes() uint64 {
	return uint64(C.xenstat_network_rbytes(n.n))
}

// RPackets returns the number of receive packets for this network
func (n *Network)RPackets() uint64 {
	return uint64(C.xenstat_network_rpackets(n.n))
}

// RErrs returns the number of receive errors for this network
func (n *Network)RErrs() uint64 {
	return uint64(C.xenstat_network_rerrs(n.n))
}

// RDrop returns the number of receive drops for this network
func (n *Network)RDrop() uint64 {
	return uint64(C.xenstat_network_rdrop(n.n))
}

// TBytes returns the number of transmit bytes for this network
func (n *Network)TBytes() uint64 {
	return uint64(C.xenstat_network_tbytes(n.n))
}

// TPackets returns the number of transmit packets for this network
func (n *Network)TPackets() uint64 {
	return uint64(C.xenstat_network_tpackets(n.n))
}

// TErrs returns the number of transmit errors for this network
func (n *Network)TErrs() uint64 {
	return uint64(C.xenstat_network_terrs(n.n))
}

// TDrop returns the number of transmit drops for this network
func (n *Network)TDrop() uint64 {
	return uint64(C.xenstat_network_tdrop(n.n))
}