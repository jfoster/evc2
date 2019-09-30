package evc2

import (
	"github.com/google/gousb"
)

func GetEVCs() (evcs []*EVC2, err error) {
	evcs, err = getEVCs(Vid, Pid)

	return
}

func getEVCs(vid, pid gousb.ID) (evcs []*EVC2, err error) {
	ctx := gousb.NewContext()

	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		return desc.Vendor == vid && desc.Product == pid
	})

	for i, device := range devs {
		evcs = append(evcs, &EVC2{
			context: ctx,
			device:  device,
		})

		for num := range device.Desc.Configs {
			config, err := device.Config(num)
			if err != nil {
				continue
			}
			defer config.Close()

			for _, ifacedesc := range config.Desc.Interfaces {
				iface, err := config.Interface(ifacedesc.Number, 0)
				if err != nil {
					continue
				}
				evcs[i].iface = iface

				for _, endpointDesc := range iface.Setting.Endpoints {
					switch endpointDesc.Direction {
					case gousb.EndpointDirectionIn:
						evcs[i].inpipe, err = iface.InEndpoint(endpointDesc.Number)
						if err != nil {
							continue
						}
						break
					case gousb.EndpointDirectionOut:
						evcs[i].outpipe, err = iface.OutEndpoint(endpointDesc.Number)
						if err != nil {
							continue
						}
					}
				}
			}
		}
	}

	return
}
