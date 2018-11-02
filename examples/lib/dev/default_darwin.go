package dev

import (
	"github.com/GreenLightning/ble"
	"github.com/GreenLightning/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
