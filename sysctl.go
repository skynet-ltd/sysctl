package sysctl

/*
 #include "query_sysctl.h"
 struct response *query_sysctl(const char *, char **);
*/
import "C"

import (
	"encoding/binary"
	"errors"
	"strconv"
	"unsafe"
)

// StringTypes ...
var StringTypes = map[string]struct{}{
	"hw.machine":    struct{}{},
	"hw.targettype": struct{}{},
	"hw.model":      struct{}{},
	"hw.cachesize":  struct{}{},

	"machdep.cpu.brand_string":           struct{}{},
	"machdep.cpu.features":               struct{}{},
	"machdep.cpu.leaf7_features":         struct{}{},
	"machdep.cpu.extfeatures":            struct{}{},
	"machdep.cpu.vendor":                 struct{}{},
	"machdep.misc.timer_queue_trace":     struct{}{},
	"machdep.xcpm.deep_idle_last_stats":  struct{}{},
	"machdep.xcpm.deep_idle_total_stats": struct{}{},
	"ktrace.configured_by":               struct{}{},

	"kern.uuid":          struct{}{},
	"user.cs_path":       struct{}{},
	"kern.ostype":        struct{}{},
	"kern.hostname":      struct{}{},
	"kern.procname":      struct{}{},
	"kern.nisdomainname": struct{}{},
	"kern.corefile":      struct{}{},
	"kern.threadname":    struct{}{},
	"kern.bootargs":      struct{}{},
	"kern.sched":         struct{}{},
	"kern.wakereason":    struct{}{},
	"kern.hibernatefile": struct{}{},
	"vm.swapfileprefix":  struct{}{},
	"net.link.generic.system.port_used.wakeuuid_not_set_last_if": struct{}{},
	"debug.swd_kext_name":  struct{}{},
	"debug.swd_delay_type": struct{}{},
}

// Call ...
func Call(name string) (string, error) {
	var e *C.char
	resp := C.query_sysctl(C.CString(name), &e)
	err := C.GoString(e)
	if err != "" {
		return "", errors.New(err)
	}

	if _, ok := StringTypes[name]; ok {
		return C.GoString(resp.buffer), nil
	}
	var n int64
	if int(resp.len) <= 4 {
		n = int64(binary.LittleEndian.Uint32(C.GoBytes(unsafe.Pointer(resp.buffer), C.int(resp.len))))
	} else {
		n = int64(binary.LittleEndian.Uint64(C.GoBytes(unsafe.Pointer(resp.buffer), C.int(resp.len))))
	}
	return strconv.FormatInt(n, 10), nil
}
