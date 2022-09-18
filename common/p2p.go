package common

import (
	"encoding/binary"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net"
	"strconv"
	"time"
)

var bufferSize = 10
var id2addr map[uint64]net.UDPAddr
var maxMapSize = 1000

type RecvPocket struct {
	Id uint64
}

type SendPocket struct {
	Ip net.IP
}

func p2pLog(s string) {
	t := time.Now()
	_, _ = fmt.Fprintf(gin.DefaultWriter, "[P2P] %v | %s \n", t.Format("2006/01/02 - 15:04:05"), s)
}

// Server Response Protocol:
// The first byte is the status byte, is an uint8 number, 0 means okay, 1 means error.
// The remaining bytes is the data to be sent.

var OkayByte uint8 = 0
var ErrorByte uint8 = 1

func sendIdAssignmentPocket(server *net.UDPConn, receiverAddr *net.UDPAddr, id uint64) {
	buffer := make([]byte, 9)
	buffer[0] = OkayByte
	binary.LittleEndian.PutUint64(buffer[1:], id)
	_, err := server.WriteToUDP(buffer, receiverAddr)
	if err != nil {
		p2pLog("error during send to " + receiverAddr.String())
		return
	}
}

func sendIdNotFoundPocket(server *net.UDPConn, receiverAddr *net.UDPAddr) {
	buffer := make([]byte, 1)
	buffer[0] = ErrorByte
	_, err := server.WriteToUDP(buffer, receiverAddr)
	if err != nil {
		p2pLog("error during send to " + receiverAddr.String())
		return
	}
}

func sendConnPockets(server *net.UDPConn, senderAddr *net.UDPAddr, receiverAddr *net.UDPAddr) {
	buffer := make([]byte, 1)
	buffer[0] = OkayByte
	buffer = append(buffer, []byte(senderAddr.String())...)
	_, err := server.WriteToUDP(buffer, receiverAddr)
	if err != nil {
		p2pLog("error during send to " + receiverAddr.String())
		return
	}
	buffer = buffer[:1]
	buffer = append(buffer, []byte(receiverAddr.String())...)
	_, err = server.WriteToUDP(buffer, senderAddr)
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
		id := binary.LittleEndian.Uint64(buffer[:8])
		if id == 0 {
			// Sender request a unique id
			if len(id2addr) > maxMapSize {
				p2pLog("too many items in id2addr, reset it")
				id2addr = make(map[uint64]net.UDPAddr)
			}
			for {
				id = rand.Uint64()
				if _, ok := id2addr[id]; ok {
					continue
				} else {
					break
				}
			}
			p2pLog("assign id " + strconv.FormatUint(id, 10) + " to " + addr.String())
			id2addr[id] = *addr
			go sendIdAssignmentPocket(server, addr, id)
		} else {
			// Receiver register with id
			p2pLog("register id " + strconv.FormatUint(id, 10) + " with " + addr.String())
			if addr2, ok := id2addr[id]; ok {
				p2pLog("send connection pockets to " + addr.String() + " & " + addr2.String())
				delete(id2addr, id)
				go sendConnPockets(server, &addr2, addr)
			} else {
				go sendIdNotFoundPocket(server, addr)
			}
		}
	}
}
