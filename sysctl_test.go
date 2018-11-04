package sysctl

import (
	"strings"
	"testing"
)

var hwQuery = "hw.machine hw.model hw.ncpu hw.byteorder hw.physmem hw.usermem hw.pagesize_compat hw.epoch hw.vectorunit hw.busfrequency_compat hw.cpufrequency_compat hw.cachelinesize_compat hw.l1icachesize_compat hw.l1dcachesize_compat hw.l2settings hw.l2cachesize_compat hw.l3settings hw.l3cachesize_compat hw.tbfrequency_compat hw.memsize hw.activecpu hw.physicalcpu hw.physicalcpu_max hw.logicalcpu hw.logicalcpu_max hw.cputype hw.cpusubtype hw.cpu64bit_capable hw.cpufamily hw.cacheconfig hw.cachesize hw.pagesize hw.pagesize32 hw.busfrequency hw.busfrequency_min hw.busfrequency_max hw.cpufrequency hw.cpufrequency_min hw.cpufrequency_max hw.cachelinesize hw.l1icachesize hw.l1dcachesize hw.l2cachesize hw.l3cachesize hw.tbfrequency hw.packages hw.optional.floatingpoint hw.optional.mmx hw.optional.sse hw.optional.sse2 hw.optional.sse3 hw.optional.supplementalsse3 hw.optional.sse4_1 hw.optional.sse4_2 hw.optional.x86_64 hw.optional.aes hw.optional.avx1_0 hw.optional.rdrand hw.optional.f16c hw.optional.enfstrg hw.optional.fma hw.optional.avx2_0 hw.optional.bmi1 hw.optional.bmi2 hw.optional.rtm hw.optional.hle hw.optional.adx hw.optional.mpx hw.optional.sgx hw.optional.avx512f hw.optional.avx512cd hw.optional.avx512dq hw.optional.avx512bw hw.optional.avx512vl hw.optional.avx512ifma hw.optional.avx512vbmi hw.targettype hw.cputhreadtype debug.intel.schedThrottleLowPriByVal"
var machdepQuery = "machdep.user_idle_level machdep.cpu.max_basic machdep.cpu.max_ext machdep.cpu.vendor machdep.cpu.brand_string machdep.cpu.family machdep.cpu.model machdep.cpu.extmodel machdep.cpu.extfamily machdep.cpu.stepping machdep.cpu.feature_bits machdep.cpu.leaf7_feature_bits machdep.cpu.extfeature_bits machdep.cpu.signature machdep.cpu.brand machdep.cpu.features machdep.cpu.leaf7_features machdep.cpu.extfeatures machdep.cpu.logical_per_package machdep.cpu.cores_per_package machdep.cpu.microcode_version machdep.cpu.processor_flag machdep.cpu.mwait.linesize_min machdep.cpu.mwait.linesize_max machdep.cpu.mwait.extensions machdep.cpu.mwait.sub_Cstates machdep.cpu.thermal.sensor machdep.cpu.thermal.dynamic_acceleration machdep.cpu.thermal.invariant_APIC_timer machdep.cpu.thermal.thresholds machdep.cpu.thermal.ACNT_MCNT machdep.cpu.thermal.core_power_limits machdep.cpu.thermal.fine_grain_clock_mod machdep.cpu.thermal.package_thermal_intr machdep.cpu.thermal.hardware_feedback machdep.cpu.thermal.energy_policy machdep.cpu.xsave.extended_state machdep.cpu.xsave.extended_state1 machdep.cpu.arch_perf.version machdep.cpu.arch_perf.number machdep.cpu.arch_perf.width machdep.cpu.arch_perf.events_number machdep.cpu.arch_perf.events machdep.cpu.arch_perf.fixed_number machdep.cpu.arch_perf.fixed_width machdep.cpu.cache.linesize machdep.cpu.cache.L2_associativity machdep.cpu.cache.size machdep.cpu.address_bits.physical machdep.cpu.address_bits.virtual machdep.cpu.core_count machdep.cpu.thread_count machdep.cpu.tsc_ccc.numerator machdep.cpu.tsc_ccc.denominator machdep.vectors.timer machdep.vectors.IPI machdep.pmap.hashwalks machdep.pmap.hashcnts machdep.pmap.hashmax machdep.pmap.kernel_text_ps machdep.pmap.kern_pv_reserve machdep.memmap.Conventional machdep.memmap.RuntimeServices machdep.memmap.ACPIReclaim machdep.memmap.ACPINVS machdep.memmap.PalCode machdep.memmap.Reserved machdep.memmap.Unusable machdep.memmap.Other machdep.tsc.frequency machdep.tsc.deep_idle_rebase machdep.tsc.at_boot machdep.tsc.rebase_abs_time machdep.tsc.nanotime.tsc_base machdep.tsc.nanotime.ns_base machdep.tsc.nanotime.scale machdep.tsc.nanotime.shift machdep.tsc.nanotime.generation machdep.misc.fast_uexc_support machdep.misc.panic_restart_timeout machdep.misc.interrupt_latency_max machdep.misc.timer_queue_trace machdep.misc.nmis machdep.xcpm.mode machdep.xcpm.hard_plimit_max_100mhz_ratio machdep.xcpm.hard_plimit_min_100mhz_ratio machdep.xcpm.soft_plimit_max_100mhz_ratio machdep.xcpm.soft_plimit_min_100mhz_ratio machdep.xcpm.tuib_plimit_max_100mhz_ratio machdep.xcpm.tuib_plimit_min_100mhz_ratio machdep.xcpm.tuib_enabled machdep.xcpm.power_source machdep.xcpm.bootplim machdep.xcpm.bootpst machdep.xcpm.tuib_ns machdep.xcpm.vectors_loaded_count machdep.xcpm.ratio_change_ratelimit_ns machdep.xcpm.ratio_changes_total machdep.xcpm.maxbusdelay machdep.xcpm.maxintdelay machdep.xcpm.mid_applications machdep.xcpm.mid_relaxations machdep.xcpm.mid_mode machdep.xcpm.mid_cst_control_limit machdep.xcpm.mid_mode_active machdep.xcpm.mbd_mode machdep.xcpm.mbd_applications machdep.xcpm.mbd_relaxations machdep.xcpm.forced_idle_ratio" +
	" machdep.xcpm.forced_idle_period machdep.xcpm.deep_idle_log machdep.xcpm.qos_txfr machdep.xcpm.deep_idle_count   machdep.xcpm.cpu_thermal_level machdep.xcpm.gpu_thermal_level machdep.xcpm.io_thermal_level machdep.xcpm.io_control_engages machdep.xcpm.io_control_disengages machdep.xcpm.io_filtered_reads machdep.xcpm.io_cst_control_enabled machdep.xcpm.ring_boost_enabled machdep.xcpm.io_epp_boost_enabled machdep.xcpm.epp_override machdep.eager_timer_evaluations machdep.eager_timer_evaluation_max machdep.x86_fp_simd_isr_uses"

func TestCall(t *testing.T) {
	hwSlice := strings.Fields(hwQuery)

	for _, hwq := range hwSlice {

		_, err := Call(hwq)
		if err != nil {
			t.Fatal(err)
		}
	}
	machdepQuerySlice := strings.Fields(machdepQuery)

	for _, hwq := range machdepQuerySlice {
		_, err := Call(hwq)
		if err != nil {
			t.Fatalf("%s %v", hwq, err)
		}
	}
}
