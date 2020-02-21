package uni

import (
	"io"

	"github.com/mikezange/unified-network-interface/builder"
)

type (
	CommandType byte
	Subsystem byte
)

const (
	cmd_POLL CommandType = iota
	cmd_SREQ
	cmd_AREQ
	cmd_SRSP
	cmd_RES0
	cmd_RES1
	cmd_RES2
	cmd_RES3
)

const (
	Sub_RES0 Subsystem = iota
	sub_SYS
	sub_MAC
	sub_NWK
	sub_AF
	sub_ZDO
	sub_SAPI
	sub_UTIL
	sub_DBG
	sub_APP
	sub_RCAF    // Remote Control Application Framework
	sub_RCN     // Remote Control Network Layer
)

const sof byte = 0xFE

type Uni struct {
	transceiver io.ReadWriter
	size uint8
	incoming chan byte
	errors chan error
}

type Frame struct {
	CommandType CommandType
	Subsystem Subsystem
	Command byte
	Payload []byte
}

func New(size uint8, transceiver io.ReadWriter) *Uni {
	uni := &Uni{
		transceiver: transceiver,
		size:        size,
		incoming:    make(chan byte),
		errors:      make(chan error),
	}

	go uni.receive()
	return uni
}

func (u *Uni) receive() {
	var buf [1]byte
	for {
		n, err := io.ReadFull(u.transceiver, buf[:])
		if n > 0 {
			u.incoming <- buf[0]
		} else if err != io.EOF {
			u.errors <- err
		}
	}
}

func (u *Uni) WriteFrame(frame *Frame) error {
	rendered := u.RenderFrame(frame)
	_, err := u.transceiver.Write(rendered)
	return err
}

func (u *Uni) RenderFrame(frame *Frame) []byte {
	cmd := ((byte(frame.CommandType << 5)) & 0xE0) | (byte(frame.Subsystem) & 0x1F)
	payloadLen := len(frame.Payload)

	build := builder.New()
	build.Byte(sof)

	if u.size == 1 {
		build.Uint8(uint8(payloadLen))
	} else {
		build.Uint16BE(uint16(payloadLen))
	}

	build.Byte(cmd).Byte(frame.Command).Bytes(frame.Payload)

	checksum := checksum(build.Make()[1:])
	build.Byte(checksum)

	return build.Make()
}

func checksum(data []byte) (b byte) {
	b = byte(0)
	for i := 0; i < len(data); i++ {
		b ^= data[i]
	}
	return
}