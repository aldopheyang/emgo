package ob

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f030x8/mmap"
)

type OB_Periph struct {
	RDP   RRDP
	USER  RUSER
	DATA0 RDATA0
	DATA1 RDATA1
	WRP0  RWRP0
	WRP1  RWRP1
}

func (p *OB_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var OB = (*OB_Periph)(unsafe.Pointer(uintptr(mmap.OB_BASE)))

type RDP uint16

func (b RDP) Field(mask RDP) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RDP) J(v int) RDP {
	return RDP(bits.Make32(v, uint32(mask)))
}

type RRDP struct{ mmio.U16 }

func (r *RRDP) Bits(mask RDP) RDP     { return RDP(r.U16.Bits(uint16(mask))) }
func (r *RRDP) StoreBits(mask, b RDP) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RRDP) SetBits(mask RDP)      { r.U16.SetBits(uint16(mask)) }
func (r *RRDP) ClearBits(mask RDP)    { r.U16.ClearBits(uint16(mask)) }
func (r *RRDP) Load() RDP             { return RDP(r.U16.Load()) }
func (r *RRDP) Store(b RDP)           { r.U16.Store(uint16(b)) }

type RMRDP struct{ mmio.UM16 }

func (rm RMRDP) Load() RDP   { return RDP(rm.UM16.Load()) }
func (rm RMRDP) Store(b RDP) { rm.UM16.Store(uint16(b)) }

func (p *OB_Periph) RDP() RMRDP {
	return RMRDP{mmio.UM16{&p.RDP.U16, uint16(RDP)}}
}

func (p *OB_Periph) nRDP() RMRDP {
	return RMRDP{mmio.UM16{&p.RDP.U16, uint16(nRDP)}}
}

type USER uint16

func (b USER) Field(mask USER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask USER) J(v int) USER {
	return USER(bits.Make32(v, uint32(mask)))
}

type RUSER struct{ mmio.U16 }

func (r *RUSER) Bits(mask USER) USER    { return USER(r.U16.Bits(uint16(mask))) }
func (r *RUSER) StoreBits(mask, b USER) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RUSER) SetBits(mask USER)      { r.U16.SetBits(uint16(mask)) }
func (r *RUSER) ClearBits(mask USER)    { r.U16.ClearBits(uint16(mask)) }
func (r *RUSER) Load() USER             { return USER(r.U16.Load()) }
func (r *RUSER) Store(b USER)           { r.U16.Store(uint16(b)) }

type RMUSER struct{ mmio.UM16 }

func (rm RMUSER) Load() USER   { return USER(rm.UM16.Load()) }
func (rm RMUSER) Store(b USER) { rm.UM16.Store(uint16(b)) }

func (p *OB_Periph) USER() RMUSER {
	return RMUSER{mmio.UM16{&p.USER.U16, uint16(USER)}}
}

func (p *OB_Periph) nUSER() RMUSER {
	return RMUSER{mmio.UM16{&p.USER.U16, uint16(nUSER)}}
}

type DATA0 uint16

func (b DATA0) Field(mask DATA0) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DATA0) J(v int) DATA0 {
	return DATA0(bits.Make32(v, uint32(mask)))
}

type RDATA0 struct{ mmio.U16 }

func (r *RDATA0) Bits(mask DATA0) DATA0   { return DATA0(r.U16.Bits(uint16(mask))) }
func (r *RDATA0) StoreBits(mask, b DATA0) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDATA0) SetBits(mask DATA0)      { r.U16.SetBits(uint16(mask)) }
func (r *RDATA0) ClearBits(mask DATA0)    { r.U16.ClearBits(uint16(mask)) }
func (r *RDATA0) Load() DATA0             { return DATA0(r.U16.Load()) }
func (r *RDATA0) Store(b DATA0)           { r.U16.Store(uint16(b)) }

type RMDATA0 struct{ mmio.UM16 }

func (rm RMDATA0) Load() DATA0   { return DATA0(rm.UM16.Load()) }
func (rm RMDATA0) Store(b DATA0) { rm.UM16.Store(uint16(b)) }

type DATA1 uint16

func (b DATA1) Field(mask DATA1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DATA1) J(v int) DATA1 {
	return DATA1(bits.Make32(v, uint32(mask)))
}

type RDATA1 struct{ mmio.U16 }

func (r *RDATA1) Bits(mask DATA1) DATA1   { return DATA1(r.U16.Bits(uint16(mask))) }
func (r *RDATA1) StoreBits(mask, b DATA1) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RDATA1) SetBits(mask DATA1)      { r.U16.SetBits(uint16(mask)) }
func (r *RDATA1) ClearBits(mask DATA1)    { r.U16.ClearBits(uint16(mask)) }
func (r *RDATA1) Load() DATA1             { return DATA1(r.U16.Load()) }
func (r *RDATA1) Store(b DATA1)           { r.U16.Store(uint16(b)) }

type RMDATA1 struct{ mmio.UM16 }

func (rm RMDATA1) Load() DATA1   { return DATA1(rm.UM16.Load()) }
func (rm RMDATA1) Store(b DATA1) { rm.UM16.Store(uint16(b)) }

type WRP0 uint16

func (b WRP0) Field(mask WRP0) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP0) J(v int) WRP0 {
	return WRP0(bits.Make32(v, uint32(mask)))
}

type RWRP0 struct{ mmio.U16 }

func (r *RWRP0) Bits(mask WRP0) WRP0    { return WRP0(r.U16.Bits(uint16(mask))) }
func (r *RWRP0) StoreBits(mask, b WRP0) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RWRP0) SetBits(mask WRP0)      { r.U16.SetBits(uint16(mask)) }
func (r *RWRP0) ClearBits(mask WRP0)    { r.U16.ClearBits(uint16(mask)) }
func (r *RWRP0) Load() WRP0             { return WRP0(r.U16.Load()) }
func (r *RWRP0) Store(b WRP0)           { r.U16.Store(uint16(b)) }

type RMWRP0 struct{ mmio.UM16 }

func (rm RMWRP0) Load() WRP0   { return WRP0(rm.UM16.Load()) }
func (rm RMWRP0) Store(b WRP0) { rm.UM16.Store(uint16(b)) }

func (p *OB_Periph) WRP0() RMWRP0 {
	return RMWRP0{mmio.UM16{&p.WRP0.U16, uint16(WRP0)}}
}

func (p *OB_Periph) nWRP0() RMWRP0 {
	return RMWRP0{mmio.UM16{&p.WRP0.U16, uint16(nWRP0)}}
}

type WRP1 uint16

func (b WRP1) Field(mask WRP1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP1) J(v int) WRP1 {
	return WRP1(bits.Make32(v, uint32(mask)))
}

type RWRP1 struct{ mmio.U16 }

func (r *RWRP1) Bits(mask WRP1) WRP1    { return WRP1(r.U16.Bits(uint16(mask))) }
func (r *RWRP1) StoreBits(mask, b WRP1) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RWRP1) SetBits(mask WRP1)      { r.U16.SetBits(uint16(mask)) }
func (r *RWRP1) ClearBits(mask WRP1)    { r.U16.ClearBits(uint16(mask)) }
func (r *RWRP1) Load() WRP1             { return WRP1(r.U16.Load()) }
func (r *RWRP1) Store(b WRP1)           { r.U16.Store(uint16(b)) }

type RMWRP1 struct{ mmio.UM16 }

func (rm RMWRP1) Load() WRP1   { return WRP1(rm.UM16.Load()) }
func (rm RMWRP1) Store(b WRP1) { rm.UM16.Store(uint16(b)) }

func (p *OB_Periph) WRP1() RMWRP1 {
	return RMWRP1{mmio.UM16{&p.WRP1.U16, uint16(WRP1)}}
}

func (p *OB_Periph) nWRP1() RMWRP1 {
	return RMWRP1{mmio.UM16{&p.WRP1.U16, uint16(nWRP1)}}
}