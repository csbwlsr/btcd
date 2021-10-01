package wire

import "io"

// MsgUnknown wireshark判定未知的报文
type MsgUnknown struct {

}

func (m *MsgUnknown) BtcDecode(reader io.Reader, u uint32, encoding MessageEncoding) error {
	return nil
}

func (m *MsgUnknown) BtcEncode(writer io.Writer, u uint32, encoding MessageEncoding) error {
	return nil
}

func (m *MsgUnknown) Command() string {
	return CmdCmpct
}

func (m *MsgUnknown) MaxPayloadLength(u uint32) uint32 {
  	return 0
}

func NewMsgUnknown() *MsgUnknown{
	return &MsgUnknown{}
}