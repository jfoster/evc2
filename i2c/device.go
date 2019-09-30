package i2c

import "encoding/xml"

type Device struct {
	XMLName     xml.Name `xml:"EVC2"`
	Application struct {
		SoftwareVersion string `xml:"SoftwareVersion"`
		FileVersion     string `xml:"FileVersion"`
	} `xml:"Application"`
	Device []struct {
		Name     string `xml:"Name"`
		Class    string `xml:"Class"`
		BusType  string `xml:"BusType"`
		BusSpeed string `xml:"BusSpeed"`
		Address  struct {
			Type  string `xml:"Type,attr"`
			Start string `xml:"Start,omitempty"`
			End   string `xml:"End,omitempty"`
		} `xml:"Address,omitempty"`
		Detect struct {
			Type     string `xml:"Type,attr"`
			Register []struct {
				Offset string `xml:"Offset"`
				Data   string `xml:"Data"`
			} `xml:"Register,omitempty"`
		} `xml:"Detect"`
		ConstantItem []struct {
			Name     string `xml:"Name"`
			Register struct {
				Type     string `xml:"Type,attr"`
				Offset   string `xml:"Offset"`
				StartBit string `xml:"StartBit,omitempty"`
				EndBit   string `xml:"EndBit,omitempty"`
				Math     struct {
					Offset string `xml:"Offset,omitempty"`
					Factor string `xml:"Factor,omitempty"`
					Unit   string `xml:"Unit,omitempty"`
				} `xml:"Math,omitempty"`
				Data []struct {
					Text string `xml:",chardata"`
					Desc string `xml:"Desc,attr"`
				} `xml:"Data,omitempty"`
				Length string `xml:"Length,omitempty"`
				Bit    string `xml:"Bit,omitempty"`
			} `xml:"Register"`
		} `xml:"Constant>Item"`
		StatusItem []struct {
			Register struct {
				Offset   string `xml:"Offset"`
				StartBit string `xml:"StartBit,omitempty"`
				EndBit   string `xml:"EndBit,omitempty"`
				Length   string `xml:"Length,omitempty"`
				Bit      string `xml:"Bit,omitempty"`
			} `xml:"Register"`
		} `xml:"Status>Item"`
		ConfigurationItem []struct {
			Name     string `xml:"Name,omitempty"`
			Register struct {
				Type   string `xml:"Type,attr"`
				Offset string `xml:"Offset"`
				Math   struct {
					Factor       string `xml:"Factor,omitempty"`
					Unit         string `xml:"Unit,omitempty"`
					Offset       string `xml:"Offset,omitempty"`
					SignBit      string `xml:"SignBit,omitempty"`
					DependFactor string `xml:"DependFactor,omitempty"`
				} `xml:"Math,omitempty"`
				Data []struct {
					Text string `xml:",chardata"`
					Desc string `xml:"Desc,attr"`
				} `xml:"Data,omitempty"`
				StartBit   string `xml:"StartBit,omitempty"`
				EndBit     string `xml:"EndBit,omitempty"`
				Length     string `xml:"Length,omitempty"`
				Bit        string `xml:"Bit,omitempty"`
				Dependance struct {
					Offset string `xml:"Offset,omitempty"`
					Length string `xml:"Length,omitempty"`
					Bit    string `xml:"Bit,omitempty"`
					Data   string `xml:"Data,omitempty"`
				} `xml:"Dependance,omitempty"`
			} `xml:"Register"`
		} `xml:"Configuration>Item"`
		MonitoringItem []struct {
			Name     string `xml:"Name"`
			Register struct {
				Type   string `xml:"Type,attr"`
				Offset string `xml:"Offset"`
				Math   struct {
					Factor       string `xml:"Factor,omitempty"`
					Offset       string `xml:"Offset,omitempty"`
					Unit         string `xml:"Unit,omitempty"`
					DependFactor string `xml:"DependFactor,omitempty"`
				} `xml:"Math,omitempty"`
				StartBit   string `xml:"StartBit,omitempty"`
				EndBit     string `xml:"EndBit,omitempty"`
				Length     string `xml:"Length,omitempty"`
				Dependance struct {
					Offset string `xml:"Offset,omitempty"`
					Length string `xml:"Length,omitempty"`
					Bit    string `xml:"Bit,omitempty"`
					Data   string `xml:"Data,omitempty"`
				} `xml:"Dependance,omitempty"`
			} `xml:"Register"`
			PreExec struct {
				Type     string `xml:"Type,attr"`
				Offset   string `xml:"Offset"`
				StartBit string `xml:"StartBit"`
				EndBit   string `xml:"EndBit"`
				Data     string `xml:"Data"`
			} `xml:"PreExec,omitempty"`
		} `xml:"Monitoring>Item"`
	} `xml:"Device"`
}
