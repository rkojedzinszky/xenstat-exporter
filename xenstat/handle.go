package xenstat

/*
#cgo pkg-config: xenlight
#cgo LDFLAGS: -L/usr/lib/xen-4.8/lib -Wl,-rpath /usr/lib/xen-4.8/lib -lxenstat
#include <xenstat.h>
*/
import "C"

// Handle mirrors xenstat_handle
type Handle struct {
	h *C.xenstat_handle
}

// Init returns a new xenstat.Handle
func Init() *Handle {
	if h := C.xenstat_init(); h != nil {
		return &Handle{h: h}
	}

	return nil
}

// Uninit closes a xenstat.Handle
func (h *Handle) Uninit() {
	C.xenstat_uninit(h.h)
}