package common

import (
	"encoding/binary"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
	"time"
)

var bufferSize = 10
var id2addr map[uint64]net.UDPAddr

type RecvPocket struct {
	Id uint64
}

type SendPocket struct {
	Ip net.IP
}

func p2pLog(s string) {
	t := time.Now()
	fmt.Fprintf(gin.DefaultWriter, "[P2P] %v | %s \n", t.Format("2006/01/02 - 15:04:05"), s)
}

func sendPockets(server *net.UDPConn, senderAddr net.UDPAddr, receiverAddr net.UDPAddr) {
	_, err := server.WriteToUDP([]byte(senderAddr.String()), &receiverAddr)
	if err != nil {
		p2pLog("error during send to " + receiverAddr.String())
		return
	}
	_, err = server.WriteToUDP([]byte(receiverAddr.String()), &senderAddr)
	if err != nil {
		p2pLog("error during send to " + senderAddr.String())
		return
	}
}

func StartP2PServer() {
	server, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: *P2PPort})
	if err != nil {
		p2pLog("failed to start p2p server")
	}
	p2pLog("start p2p server on " + server.LocalAddr().String())
	buffer := make([]byte, bufferSize)
	id2addr = make(map[uint64]net.UDPAddr)
	for {
		_, addr, err := server.ReadFromUDP(buffer)
		if err != nil {
			p2pLog("error during read from udp: " + err.Error())
			continue
		}
		id := binary.BigEndian.Uint64(buffer[:8])
		if addr2, ok := id2addr[id]; ok {
			p2pLog("send pocket to " + addr.String() + " & " + addr2.String())
			delete(id2addr, id)
			go sendPockets(server, addr2, *addr)
		} else {
			p2pLog("register id " + strconv.FormatUint(id, 10) + " with " + addr.String())
			id2addr[id] = *addr
		}
	}
}
