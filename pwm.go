package wPi

// #cgo LDFLAGS: -lwiringPi -lpthread
// #include<wiringPi.h>
// #include<softPwm.h>
import "C"

import (
	"fmt"
)

/*SoftPwm is used to control a pin with PWM(Pulse Width Modulation).
It is a software implementation and may cause Raspberry Pi to overload.*/
type SoftPwm struct {
	pin  int
	rang int
}

/*NewSoftPwm creates SoftPwm struct and returns.
Range and Pin number has to be provided*/
func NewSoftPwm(pin, initialValue, pwmRange int) SoftPwm {
	C.softPwmCreate(C.int(pin), C.int(initialValue), C.int(pwmRange))
	return SoftPwm{pin: pin, rang: pwmRange}
}

/*Write writes the Pwm value to the pwm pin setup earlier.*/
func (spwm SoftPwm) Write(value int) error {
	if value < 0 || value > spwm.rang {
		return fmt.Errorf("Please provide PWM range between 0 and %d", spwm.rang)
	}
	C.softPwmWrite(C.int(spwm.pin), C.int(value))
	return nil
}
