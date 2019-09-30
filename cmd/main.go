package main

import (
	"fmt"
	"log"

	"github.com/jfoster/evc2"
	"github.com/jfoster/evc2/i2c"
)

func main() {
	evcs, err := evc2.GetEVCs()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(len(evcs))
	for _, evc := range evcs {
		log.Println("IsConnected", evc.IsConnected())
		log.Println("IsDFU", evc.IsDFU())
		log.Println("DeviceUID", evc.DeviceUID())
		log.Println("FirmwareVersion", evc.FirmwareVersion())
		log.Println("DeviceID ", fmt.Sprintf("%X", evc.DeviceID()))

		defer evc.Close()
	}

	devices, err := i2c.GetDevices()

	for i, device := range devices {
		fmt.Println(i, device)
	}
}
