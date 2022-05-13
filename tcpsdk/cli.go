package tcpsdk

import (
	"io"
	"log"

	"github.com/team-tech-lmh/zenith/sockets"
)

type Client struct {
	sock *sockets.Socket
}

func NewClient(ip string, port int) (*Client, error) {
	soc := sockets.Socket{
		IPAddr: ip,
		Port:   port,
	}
	if err := soc.Connect(); nil != err {
		return nil, err
	}
	return &Client{
		sock: &soc,
	}, nil
}

func (cli *Client) SendCmd(cmd Cmd) error {
	if l, err := cli.sock.Conn.Write(cmd.ToBytes()); nil != err {
		return err
	} else {
		log.Printf("send cmd, length: %v", l)
	}
	return nil
}

func (cli *Client) ReceiveCmd() (*Cmd, error) {
	header, err := cli.receiveCmdHearder()
	if nil != err {
		log.Printf("read cmd failed when read header %v\n", err)
		return nil, err
	}

	cmd, err := cli.receiveCmdData(header)
	if nil != err {
		log.Printf("read cmd failed when read data %v\n", err)
		return nil, err
	}
	return cmd, nil
}

func (cli *Client) ReceiveRaw(l int) ([]byte, error) {
	needL := l
	buf := make([]byte, needL)
	if _, err := cli.sock.Conn.Read(buf); nil != err {
		if io.EOF == err {
			return nil, nil
		}
		return nil, err
	} else {
		return buf, nil
	}
}

func (cli *Client) receiveCmdHearder() (*CmdHeader, error) {
	relBuf := []byte{}
	for {
		if len(relBuf) >= cmdHederLen {
			break
		}
		buf, err := cli.ReceiveRaw(cmdHederLen - len(relBuf))
		if nil != err {
			log.Printf("read header failed with buf not full : %v\n", err)
			return nil, err
		}
		relBuf = append(relBuf, buf...)
	}
	header := CmdHeaderFromBuf(relBuf)
	return &header, nil
}

func (cli *Client) receiveCmdData(header *CmdHeader) (*Cmd, error) {
	if header.PackageLen <= 0 {
		return &Cmd{
			CmdHeader: *header,
		}, nil
	}

	relBuf := []byte{}
	for {
		if len(relBuf) >= int(header.PackageLen) {
			break
		}
		buf, err := cli.ReceiveRaw(int(header.PackageLen) - len(relBuf))
		if nil != err {
			log.Printf("read header failed with buf not full : %v\n", err)
			return nil, err
		}
		relBuf = append(relBuf, buf...)
	}
	return &Cmd{
		CmdHeader: *header,
		Data:      relBuf,
	}, nil
}
