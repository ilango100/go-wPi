// Package wPi provides a wrapper library for wiringPi library
package wPi

// #cgo LDFLAGS: -lwiringPi
// #include<wiringPi.h>
import "C"

// wiringPi Modes
const (
	WPI_MODE_PINS int = iota
	WPI_MODE_GPIO
	WPI_MODE_GPIO_SYS
	WPI_MODE_PHYS
	WPI_MODE_PIFACE
	WPI_MODE_UNINITIALISED int = -1
)

// Pin Modes
const (
	INPUT int = iota
	OUTPUT
	PWM_OUTPUT
	GPIO_CLOCK
	SOFT_PWM_OUTPUT
	SOFT_TONE_OUTPUT
	PWM_TONE_OUTPUT
)
