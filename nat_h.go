package main

import (
    "net"
    "encoding/binary"
    "time"
)

type Packet []byte
type OpCode byte
type ResultCode byte

const (
    ExtIP OpCode = 0
    MapUDP OpCode = 1
    MapTCP OpCode = 2
)
const (
    Success ResultCode = 0
    BadVers ResultCode = 1
    Refused ResultCode = 2
    NetFail ResultCode = 3
    OutOfRes ResultCode = 4
    BadOP ResultCode = 5
)

func NewPacket() Packet {
    p := make(Packet, 12)
    return p
}
func (p *Packet) Truncate() {
    if p.GetOpCode() == ExtIP {
        *p = make(Packet, 2)
    }
}
func (p Packet) SetOpCode(o OpCode) {
    p[1] = byte(o)
}
func (p Packet) GetOpCode() OpCode {
    return OpCode(p[1])
}
func (p Packet) GetResultCode() ResultCode {
    return ResultCode(p[3])
}
func (p Packet) GetSecs() time.Time {
    s, _ := binary.Uvarint(p[4:8])
    return time.Unix(int64(s), 0)
}
func (p Packet) SetInternalPort(port uint16) {
    buf := make([]byte, 2)
    binary.PutUvarint(buf, uint64(port))
    copy(p[4:6], buf)
}
func (p Packet) GetInternalPort() []byte {
    return p[4:6]
}
func (p Packet) SetExternalPort(port uint16) {
    buf := make([]byte, 2)
    binary.PutUvarint(buf, uint64(port))
    copy(p[6:8], buf)
}
func (p Packet) GetExternalPort() []byte {
    return p[6:8]
}
func (p Packet) SetTTL(ttl uint32) {
    buf := make([]byte, 4)
    binary.PutUvarint(buf, uint64(ttl))
    copy(p[8:12], buf)
}
func (p Packet) GetTTL() uint32 {
    buf, _ := binary.Uvarint(p[8:12])
    return uint32(buf)
}
func (p Packet) GetExternalIP() net.IP {
    return net.IP(p[8:12])
}