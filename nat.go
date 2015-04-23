package main

import (
    "fmt"
    "net"
)

func main() {
    fmt.Println("Getting NAT external IP")
    p := NewPacket()
    p.SetOpCode(MapTCP)
    p.SetInternalPort(1488)
    p.SetExternalPort(1488)
    p.SetTTL(60)
    //p.Truncate()
    //p[0] = 1
    uc, err := net.Dial("udp", "192.168.1.1:5351")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(p)
    _, err = uc.Write(p)
    if err != nil {
        fmt.Println(err)
        return
    }
    //rp := make([]byte, 12)
    _, err = uc.Read(p)
    fmt.Println(p.GetResultCode())
    fmt.Println(p.GetSecs())
    fmt.Println(p)
}
