package bluez

import (
	"log"

	"tinygo.org/x/bluetooth"
)

func ScanAll() {
	adapter := bluetooth.NewAdapter("hci1")

	err := adapter.Enable()
	if err != nil {
		panic("Failed to enable BLE adapter.")
	}

	err = adapter.Scan(onScan)
	if err != nil {
		panic("Failed to register scan callback")
	}
}

func onScan(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
	log.Println("found device", device.Address.String(), device.RSSI, device.LocalName())
}
