package sockets

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Socket struct {
	IPAddr string
	Port   int

	Conn *net.TCPConn
}

func (opt Socket) ipAddrStr() string {
	return fmt.Sprintf("%v:%v", opt.IPAddr, opt.Port)
}

func (opt *Socket) Connect() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", opt.ipAddrStr())
	if nil != err {
		return nil
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		return err
	}
	opt.Conn = conn
	if err := conn.SetKeepAlive(true); nil != err {
		log.Printf("tcp conn set keep alive failed : %v\n", err)
	}
	if err := conn.SetKeepAlivePeriod(time.Minute); nil != err {
		log.Printf("tcp conn set keep alive failed : %v\n", err)
	}

	if err := conn.SetNoDelay(true); nil != err {
		log.Printf("tcp conn set no delay failed : %v\n", err)
	}
	if err := conn.SetReadBuffer(1024); nil != err {
		log.Printf("tcp conn set read buffer failed : %v\n", err)
	}
	return nil
}

func (opt *Socket) Close() {
	if err := opt.Conn.CloseRead(); nil != err {
		log.Printf("tcp conn closed read failed %v\n", err)
	}
	if err := opt.Conn.CloseWrite(); nil != err {
		log.Printf("tcp conn closed write failed %v\n", err)
	}
	if err := opt.Conn.Close(); nil != err {
		log.Printf("tcp conn closed failed %v\n", err)
	}
}
