package xenstat

/*
#include <xenstat.h>
*/
import "C"

// VCPU mirrors xenstat_vcpu
type VCPU struct {
	Idx uint
	v *C.xenstat_vcpu
}

// Online returns wether a VCPU is online
func (v *VCPU) Online() uint {
	return uint(C.xenstat_vcpu_online(v.v))
}

// Usage returns CPU usage of a VCPU
func (v *VCPU) Usage() uint64 {
	return uint64(C.xenstat_vcpu_ns(v.v))
}