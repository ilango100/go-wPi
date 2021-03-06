package wPi

import (
	"time"
)

// Stepper is used to operate a stepper motor
type Stepper struct {
	a, b, c, d, steps, step int
	delay                   time.Duration
}

// NewStepper creates a Stepper type using the supplied pin values.
// Initial delay value is 50 milliseconds
func NewStepper(steps, a, b, c, d int) *Stepper {
	s := &Stepper{steps: steps, a: a, b: b, c: c, d: d, delay: 50 * time.Millisecond}
	s.init()
	return s
}

// Initialises the stepper struct
func (s *Stepper) init() {
	PinMode(s.a, Output)
	PinMode(s.b, Output)
	PinMode(s.c, Output)
	PinMode(s.d, Output)
	s.write(10)
}

// Writes the step value to the stepper
// The order:
// 10, 6, 5, 9
func (s *Stepper) write(val int) {
	DigitalWrite(s.a, (val>>3)&1)
	DigitalWrite(s.b, (val>>2)&1)
	DigitalWrite(s.c, (val>>1)&1)
	DigitalWrite(s.d, val&1)
	s.step = val
}

// step steps the motor by one step
func (s *Stepper) ostep() {
	switch s.step {
	case 10:
		s.write(6)
	case 6:
		s.write(5)
	case 5:
		s.write(9)
	case 9:
		fallthrough
	default:
		s.write(10)
	}
}

// rstep steps the motor by one step in reverse direction
func (s *Stepper) rstep() {
	switch s.step {
	case 10:
		s.write(9)
	case 9:
		s.write(5)
	case 5:
		s.write(6)
	case 6:
		fallthrough
	default:
		s.write(10)
	}
}

// Step steps the motor by t steps.
// You can provide negative values too for opposite direction stepping
func (s *Stepper) Step(t int) {
	if t > 0 {
		for ; t > 0; t-- {
			s.ostep()
			time.Sleep(s.delay)
		}
	} else {
		t = -t
		for ; t > 0; t-- {
			s.rstep()
			time.Sleep(s.delay)
		}
	}
}

// SetDelay sets the delay time between each step
func (s *Stepper) SetDelay(d time.Duration) {
	s.delay = d
}

// SetSpeed automatically calculates the delay time using the rpm value and sets it
func (s *Stepper) SetSpeed(rpm int) {
	s.delay = time.Duration((60*1000*1000)/(rpm*s.steps)) * time.Microsecond
}
