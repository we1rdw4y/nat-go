package nat

import (
    "net"
    "encoding/binary"
    "time"
)

type Packet []byte
type RqPacket Packet
type AnsPacket Packet
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

func NewRqPacket() RqPacket {
    p := make(RqPacket, 12)
    return p
}
func NewAnsPacket() AnsPacket {
    p := make(AnsPacket, 16)
    return p
}
func (p RqPacket) SetOpCode(o OpCode) {
    p[1] = byte(o)
}
func (p RqPacket) OpCode() OpCode {
    return OpCode(p[1])
}
func (p AnsPacket) SetOpCode(o OpCode) {
    p[1] = byte(o + 128)
}
func (p AnsPacket) OpCode() OpCode {
    return OpCode(p[1] - 128)
}
func (p AnsPacket) SetResultCode(c ResultCode) {
    p[3] = byte(c)
}
func (p AnsPacket) ResultCode() ResultCode {
    return ResultCode(p[3])
}
func (p AnsPacket) SetSecs(t time.Time) {
    buf := make([]byte, 4)
    binary.PutUvarint(buf, uint64(uint32(t.Unix())))
    copy(p[4:8], buf)
}
func (p AnsPacket) Secs() time.Time {
    s, _ := binary.Uvarint(p[4:8])
    return time.Unix(int64(s), 0)
}
func (p RqPacket) SetInternalPort(port uint16) {
    buf := make([]byte, 2)
    binary.PutUvarint(buf, uint64(port))
    copy(p[4:6], buf)
}
func (p RqPacket) InternalPort() uint16 {
    port, _ := binary.Uvarint(p[4:6])
    return uint16(port)
}
func (p RqPacket) SetExternalPort(port uint16) {
    buf := make([]byte, 2)
    binary.PutUvarint(buf, uint64(port))
    copy(p[6:8], buf)
}
func (p RqPacket) ExternalPort() uint16 {
    port, _ := binary.Uvarint(p[6:8])
    return uint16(port)
}
func (p AnsPacket) SetInternalPort(port uint16) {
    buf := make([]byte, 2)
    binary.PutUvarint(buf, uint64(port))
    copy(p[8:10], buf)
}
func (p AnsPacket) InternalPort() uint16 {
    port, _ := binary.Uvarint(p[8:10])
    return uint16(port)
}
func (p AnsPacket) SetExternalPort(port uint16) {
    buf := make([]byte, 2)
    binary.PutUvarint(buf, uint64(port))
    copy(p[10:12], buf)
}
func (p AnsPacket) ExternalPort() uint16 {
    port, _ := binary.Uvarint(p[10:12])
    return uint16(port)
}
func (p RqPacket) SetTTL(ttl uint32) {
    buf := make([]byte, 4)
    binary.PutUvarint(buf, uint64(ttl))
    copy(p[8:12], buf)
}
func (p RqPacket) TTL() uint32 {
    buf, _ := binary.Uvarint(p[8:12])
    return uint32(buf)
}
func (p AnsPacket) SetTTL(ttl uint32) {
    buf := make([]byte, 4)
    binary.PutUvarint(buf, uint64(ttl))
    copy(p[12:16], buf)
}
func (p AnsPacket) TTL() uint32 {
    buf, _ := binary.Uvarint(p[12:16])
    return uint32(buf)
}
func (p AnsPacket) SetExternalIP(ip net.IP) {
    copy(p[8:12], ip.To4())
}
func (p AnsPacket) ExternalIP() net.IP {
    return net.IP(p[8:12])
}