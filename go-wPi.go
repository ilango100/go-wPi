// Package wPi provides as simple wrapper library for wiringPi library.
//
//Note:
// Any of the setup functions WPi, GPio, Phys, Sys should be called before any other operations
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

// WPi sets up wiringPi using wiringPI pin numbering
func WPi() {
	C.wiringPiSetup()
}

// Gpio sets up wiringPi using GPIO pin numbering
func Gpio() {
	C.wiringPiSetupGpio()
}

// Phys sets up wiringPi using Physical pin numbering
func Phys() {
	C.wiringPiSetupPhys()
}

// Sys sets up wiringPi using GPIO numbering with sysfs interface
func Sys() {
	C.wiringPiSetupSys()
}

// PinMode sets the mode of pin, can be Input, Output, PwmOutput or GpioClock
// Read documentation at wiringPi.com for information
func PinMode(pin, mode int) {
	C.pinMode(C.int(pin), C.int(mode))
}

// DigitalWrite writes the value on to the pin
// 0 - LOW, Any other value - HIGH
func DigitalWrite(pin, value int) {
	C.digitalWrite(C.int(pin), C.int(value))
}

// DigitalRead reads the input from pin
func DigitalRead(pin int) bool {
	return C.digitalRead(C.int(pin)) == C.TRUE
}
