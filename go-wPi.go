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
