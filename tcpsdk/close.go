package tcpsdk

import "encoding/json"

type IOChannel int

const (
	IOChannel1 IOChannel = iota
	IOChannel2
)

func (cli *Client) OpenBarrier() (*Cmd, error) {
	return cli.sendIOCtl(IOChannel1)
}
func (cli *Client) CloseBarrier() (*Cmd, error) {
	return cli.sendIOCtl(IOChannel2)
}

func (cli *Client) sendIOCtl(ioChannel IOChannel) (*Cmd, error) {
	cmdMap := map[string]interface{}{
		"cmd":   "ioctl_resp",
		"io":    ioChannel,
		"value": 2,
	}
	cmdBuf, err := json.Marshal(cmdMap)
	if nil != err {
		return nil, err
	}
	cmd := NewDataCmd(cmdBuf, 0, CmdTypeData)
	if err := cli.SendCmd(cmd); nil != err {
		return nil, err
	}

	if c, err := cli.ReceiveCmd(); nil != err {
		return nil, err
	} else {
		return c, nil
	}
}
