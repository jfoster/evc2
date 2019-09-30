package evc2

import (
	"encoding/binary"
	"fmt"

	"github.com/google/gousb"
)

const (
	Manufacturer = "ElmorLabs"
	Name         = "EVC 2.0"

	Vid    = gousb.ID(0x0483)
	Pid    = gousb.ID(0x5740)
	PidDFU = gousb.ID(0xDF11)
)

type EVC2 struct {
	context *gousb.Context
	device  *gousb.Device
	iface   *gousb.Interface
	inpipe  *gousb.InEndpoint
	outpipe *gousb.OutEndpoint

	// DeviceUID STM32UID

	done func()
}

func (e *EVC2) Open() (err error) {
	return e.openDevice(Vid, Pid)
}

func (e *EVC2) openDevice(vid gousb.ID, pid gousb.ID) (err error) {
	if e.context == nil {
		e.context = gousb.NewContext()
	}

	if e.device == nil {
		e.device, err = e.context.OpenDeviceWithVIDPID(vid, pid)
		if err != nil {
			return
		}
	}

	e.iface, e.done, err = e.device.DefaultInterface()

	e.inpipe, err = e.iface.InEndpoint(2)
	if err != nil {
		fmt.Println(err)
	}
	if e.inpipe != nil {
		fmt.Println(e.inpipe.Desc.Address)
	}

	return
}

func (e *EVC2) Close() (err error) {
	// if e.inpipe != nil {
	// 	e.Close()
	// }
	// if e.outpipe != nil {
	// 	e.Close()
	// }
	if e.iface != nil {
		e.done()
		e.iface.Close()
	}
	if e.device != nil {
		e.device.Close()
	}
	if e.context != nil {
		err = e.context.Close()
	}
	return
}

func (e *EVC2) WriteCommand(request byte, value uint16) (err error) {
	_, err = e.device.Control(gousb.ControlOut|gousb.ControlClass|gousb.ControlInterface, request, value, 0, nil)
	return
}

func (e *EVC2) Read(buffer []byte) (n int, err error) {
	return e.inpipe.Read(buffer)
}

func (e *EVC2) IsConnected() (connected bool) {
	if e.device != nil {
		if man, err := e.device.Manufacturer(); man == Manufacturer && err == nil {
			connected = true
		}
	}
	return
}

func (e *EVC2) IsDFU() (dfu bool) {
	if e.IsConnected() && e.ProductID() == PidDFU {
		dfu = true
	}
	return
}

func (e *EVC2) SerialNumber() (sn string) {
	if e.IsConnected() {
		var err error
		sn, err = e.device.SerialNumber()
		if err != nil {
			sn = fmt.Sprintln(err)
		}
	}
	return
}

func (e *EVC2) VendorID() (vid gousb.ID) {
	if e.IsConnected() {
		vid = e.device.Desc.Vendor
	}
	return
}

func (e *EVC2) ProductID() (pid gousb.ID) {
	if e.IsConnected() {
		pid = e.device.Desc.Product
	}
	return
}

func (e *EVC2) DeviceID() (id uint16) {
	if !e.IsConnected() {
		return
	}

	e.WriteCommand(1, 0)

	buffer := make([]byte, 3)
	if n, err := e.Read(buffer); err == nil && n == len(buffer) {
		id = binary.LittleEndian.Uint16(buffer[1:])
	}

	return
}

func (e *EVC2) HardwareVersion() (hw uint8) {
	if !e.IsConnected() {
		return
	}

	e.WriteCommand(3, 0)

	buffer := make([]byte, 2)
	if n, err := e.Read(buffer); err == nil && n == len(buffer) {
		hw = buffer[1]
	}
	return
}

func (e *EVC2) FirmwareVersion() (fw uint16) {
	if !e.IsConnected() {
		return
	}

	if err := e.WriteCommand(4, 0); err != nil {

	}

	buffer := make([]byte, 3)
	if n, err := e.Read(buffer); err == nil && n == len(buffer) {
		fw = binary.LittleEndian.Uint16(buffer[1:])
	}

	return
}

func (e *EVC2) DeviceUID() (uid STM32UID) {
	if !e.IsConnected() {
		return
	}

	if err := e.WriteCommand(5, 0); err != nil {

	}

	buffer := make([]byte, 13)
	if n, err := e.Read(buffer); err == nil && n == len(buffer) {
		uid.Unpack(buffer[1:])
	}
	return
}
