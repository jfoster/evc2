package i2c

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	i2cDevicesDir = "I2C_DEVICES"
)

type Devices []Device

func GetDevices() (devices Devices, err error) {
	devices, err = getDevices()
	return
}

func getDevices() (devices Devices, err error) {
	var files []string

	err = filepath.Walk(i2cDevicesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || strings.ToLower(filepath.Ext(path)) != ".xml" {
			return nil
		}

		files = append(files, path)
		return nil
	})

	devices = make(Devices, len(files))
	for i, file := range files {
		var data []byte

		data, err = ioutil.ReadFile(file)
		if err != nil {
			return
		}

		var device Device

		err = xml.Unmarshal(data, &device)
		if err != nil {
			return
		}

		// devices = append(devices, device)
		devices[i] = device
	}
	return
}
