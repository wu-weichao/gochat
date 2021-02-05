package message

import (
	"encoding/binary"
	"io"
)

var MsgHeader = [2]byte{'v', 1}
var VerStart = 0
var VerEnd = 2
var LenStart = VerEnd
var LenEnd = LenStart + 4

// message package
type Package struct {
	Ver [2]byte // 需指定长度
	Len int32   // 需指定类型
	Msg []byte
}

func (p *Package) Pack(w io.Writer) (err error) {
	err = binary.Write(w, binary.BigEndian, p.Ver)
	err = binary.Write(w, binary.BigEndian, p.Len)
	err = binary.Write(w, binary.BigEndian, p.Msg)
	return
}

func (p *Package) Unpack(r io.Reader) (err error) {
	err = binary.Read(r, binary.BigEndian, &p.Ver)
	err = binary.Read(r, binary.BigEndian, &p.Len)
	p.Msg = make([]byte, p.Len)
	err = binary.Read(r, binary.BigEndian, &p.Msg)
	return
}

// 协议 [2][4][...]
func NewPackageV1(msg []byte) *Package {
	return &Package{
		Ver: MsgHeader,
		Len: int32(len(msg)),
		Msg: msg,
	}
}
