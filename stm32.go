package evc2

import (
	"encoding/binary"
	"fmt"

	"github.com/go-restruct/restruct"
)

type STM32UID struct {
	X   uint16
	Y   uint16
	WAF byte
	LOT [7]byte
}

func (stm *STM32UID) Unpack(bytes []byte) {
	restruct.Unpack(bytes, binary.LittleEndian, &*stm)
}

func (stm STM32UID) String() (str string) {
	for i := range stm.LOT {
		str += fmt.Sprintf("%02X", stm.LOT[i])
	}
	str += fmt.Sprintf("%02X%04X%04X", stm.WAF, stm.Y, stm.X)
	return
}
