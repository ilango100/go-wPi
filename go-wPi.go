// Package wPi provides as simple wrapper library for wiringPi library
package wPi

// #cgo LDFLAGS: -lwiringPi
// #include<wiringPi.h>
import "C"

// Pin Modes
const (
	Input int = iota
	Output
	PwmOutput
	GpioClock
	SoftPwmOutput
	SoftToneOutput
	PwmToneOutput
)

// WiringPi setup using wiringPI pin numbering
func WiringPiSetup() {
	C.wiringPiSetup()
}

// WiringPi setup using GPIO pin numbering
func WiringPiSetupGpio() {
	C.wiringPiSetupGpio()
}

// WiringPi setup using Physical pin numbering
func WiringPiSetupPhys() {
	C.wiringPiSetupPhys()
}

// WiringPi setup using GPIO numbering with sysfs interface
func WiringPiSetupSys() {
	C.wiringPiSetupSys()
}
