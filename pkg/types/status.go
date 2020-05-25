package types

import (
	"fmt"
	"sync"
)

type Status struct {
	Settings        Settings `json:"settings"`
	BatteryLevel    int      `json:"battery_level"`
	FirmwareVersion string   `json:"firmware_version"`
}

func (s *Status) PrettyPrint() {
	fmt.Println("Firmware   :", s.FirmwareVersion)
	fmt.Println("Battery    :", s.BatteryLevel)
	fmt.Println("Settings")
	fmt.Println(" Power     :", s.Settings.PowerLevel)
	fmt.Println(" Mode      :", s.Settings.Mode)
	fmt.Println(" Channels")
	fmt.Println("  A        :", s.Settings.Channels.A)
	fmt.Println("  B        :", s.Settings.Channels.B)
	fmt.Println("  Linked   :", s.Settings.Channels.Linked)
	fmt.Println(" Parameters")
	fmt.Println("  C        :", s.Settings.Parameters.C)
	fmt.Println("  D        :", s.Settings.Parameters.D)
}

type LockedStatus struct {
	Status Status
	Mutex  sync.Mutex
}
