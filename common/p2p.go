package common

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"strconv"
)

var bufferSize = 10
var id2addr map[uint64]net.UDPAddr
var maxMapSize = 1000

type RecvPacket struct {
	Id uint64
}

type SendPacket struct {
	Ip net.IP
}

// Server Response Protocol:
// The first byte is the status byte, is an uint8 number, 0 means okay, 1 means error.
// The remaining bytes is the data to be sent.

var OkayByte uint8 = 0
var ErrorByte uint8 = 1

func sendIdAssignmentPacket(server *net.UDPConn, receiverAddr *net.UDPAddr, id uint64) {
	buffer := make([]byte, 9)
	buffer[0] = OkayByte
	binary.LittleEndian.PutUint64(buffer[1:], id)
	_, err := server.WriteToUDP(buffer, receiverAddr)
	if err != nil {
		P2pLog("error during send to " + receiverAddr.String())
		return
	}
}

func sendIdNotFoundPacket(server *net.UDPConn, receiverAddr *net.UDPAddr) {
	buffer := make([]byte, 1)
	buffer[0] = ErrorByte
	_, err := server.WriteToUDP(buffer, receiverAddr)
	if err != nil {
		P2pLog("error during send to " + receiverAddr.String())
		return
	}
}

func sendConnPackets(server *net.UDPConn, senderAddr *net.UDPAddr, receiverAddr *net.UDPAddr) {
	buffer := make([]byte, 1)
	buffer[0] = OkayByte
	buffer = append(buffer, []byte(senderAddr.String())...)
	_, err := server.WriteToUDP(buffer, receiverAddr)
	if err != nil {
		P2pLog("error during send to " + receiverAddr.String())
		return
	}
	buffer = buffer[:1]
	buffer = append(buffer, []byte(receiverAddr.String())...)
	_, err = server.WriteToUDP(buffer, senderAddr)
	if err != nil {
		P2pLog("error during send to " + senderAddr.String())
		return
	}
}

func StartP2PServer() {
	server, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: *P2PPort})
	if err != nil {
		P2pLog("failed to start p2p server")
	}
	P2pLog(fmt.Sprintf("p2p server started at port %d", *P2PPort))
	buffer := make([]byte, bufferSize)
	id2addr = make(map[uint64]net.UDPAddr)
	for {
		_, addr, err := server.ReadFromUDP(buffer)
		if err != nil {
			P2pLog("error during read from udp: " + err.Error())
			continue
		}
		id := binary.LittleEndian.Uint64(buffer[:8])
		if id == 0 {
			// Sender request a unique id
			if len(id2addr) > maxMapSize {
				P2pLog("too many items in id2addr, reset it")
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
			P2pLog("assign id " + strconv.FormatUint(id, 10) + " to " + addr.String())
			id2addr[id] = *addr
			go sendIdAssignmentPacket(server, addr, id)
		} else {
			// Receiver register with id
			P2pLog("register id " + strconv.FormatUint(id, 10) + " with " + addr.String())
			if addr2, ok := id2addr[id]; ok {
				P2pLog("send connection Packets to " + addr.String() + " & " + addr2.String())
				delete(id2addr, id)
				go sendConnPackets(server, &addr2, addr)
			} else {
				go sendIdNotFoundPacket(server, addr)
			}
		}
	}
}
