package timer

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"
)

type Timer struct{}

func (p *Timer) r(n uint) *mmio.U32 {
	return &(*[3]mmio.U32)(unsafe.Pointer(p))[n]
}

var (
	TIM1 = (*Timer)(unsafe.Pointer(uintptr(0x40010000)))
	TIM2 = (*Timer)(unsafe.Pointer(uintptr(0x40000000)))
	TIM3 = (*Timer)(unsafe.Pointer(uintptr(0x40000400)))
	TIM4 = (*Timer)(unsafe.Pointer(uintptr(0x40000800)))
)


type CR1_Bits uint32
type CR1_Field uint16

func (p *Timer) CEN() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(CEN)} }
func (p *Timer) UDIS() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(UDIS)} }
func (p *Timer) URS() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(URS)} }
func (p *Timer) OPM() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(OPM)} }
func (p *Timer) DIR() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(DIR)} }
func (p *Timer) CMS() mmio.Field32 {return mmio.Field32{p.r(0), uint16(CMS)} }
func (p *Timer) ARPE() mmio.Bits32 {return mmio.Bits32{p.r(0), uint32(ARPE)} }
func (p *Timer) CKD() mmio.Field32 {return mmio.Field32{p.r(0), uint16(CKD)} }

func (p *Timer) CR1_Load() CR1_Bits              { return CR1_Bits(p.r(0).Load()) }
func (p *Timer) CR1_Store(b CR1_Bits)            { p.r(0).Store(uint32(b)) }
func (p *Timer) CR1_Field(f CR1_Field) int       { return p.r(0).Field(uint16(f)) }
func (p *Timer) CR1_SetField(f CR1_Field, v int) { p.r(0).SetField(uint16(f), v) }

type CR2_Bits uint32
type CR2_Field uint16

func (p *Timer) CCPC() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(CCPC)} }
func (p *Timer) CCUS() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(CCUS)} }
func (p *Timer) CCDS() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(CCDS)} }
func (p *Timer) MMS() mmio.Field32 {return mmio.Field32{p.r(1), uint16(MMS)} }
func (p *Timer) TI1S() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(TI1S)} }
func (p *Timer) OIS1() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS1)} }
func (p *Timer) OIS1N() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS1N)} }
func (p *Timer) OIS2() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS2)} }
func (p *Timer) OIS2N() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS2N)} }
func (p *Timer) OIS3() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS3)} }
func (p *Timer) OIS3N() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS3N)} }
func (p *Timer) OIS4() mmio.Bits32 {return mmio.Bits32{p.r(1), uint32(OIS4)} }

func (p *Timer) CR2_Load() CR2_Bits              { return CR2_Bits(p.r(1).Load()) }
func (p *Timer) CR2_Store(b CR2_Bits)            { p.r(1).Store(uint32(b)) }
func (p *Timer) CR2_Field(f CR2_Field) int       { return p.r(1).Field(uint16(f)) }
func (p *Timer) CR2_SetField(f CR2_Field, v int) { p.r(1).SetField(uint16(f), v) }

type SMCR_Bits uint32
type SMCR_Field uint16

func (p *Timer) SMS() mmio.Field32 {return mmio.Field32{p.r(2), uint16(SMS)} }
func (p *Timer) TS() mmio.Field32 {return mmio.Field32{p.r(2), uint16(TS)} }
func (p *Timer) MSM() mmio.Bits32 {return mmio.Bits32{p.r(2), uint32(MSM)} }
func (p *Timer) ETF() mmio.Field32 {return mmio.Field32{p.r(2), uint16(ETF)} }
func (p *Timer) ETPS() mmio.Field32 {return mmio.Field32{p.r(2), uint16(ETPS)} }
func (p *Timer) ECE() mmio.Bits32 {return mmio.Bits32{p.r(2), uint32(ECE)} }
func (p *Timer) ETP() mmio.Bits32 {return mmio.Bits32{p.r(2), uint32(ETP)} }

func (p *Timer) SMCR_Load() SMCR_Bits              { return SMCR_Bits(p.r(2).Load()) }
func (p *Timer) SMCR_Store(b SMCR_Bits)            { p.r(2).Store(uint32(b)) }
func (p *Timer) SMCR_Field(f SMCR_Field) int       { return p.r(2).Field(uint16(f)) }
func (p *Timer) SMCR_SetField(f SMCR_Field, v int) { p.r(2).SetField(uint16(f), v) }