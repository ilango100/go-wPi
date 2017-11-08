package wPi

// Stepper is used to operate a stepper motor
type Stepper struct {
	a, b, c, d, steps, step int
}

// NewStepper creates a Stepper type using the supplied pin values
func NewStepper(steps, a, b, c, d int) Stepper {
	s := Stepper{steps, a, b, c, d}
	s.init()
	return s
}

// Initialises the stepper struct
func (s Stepper) init() {
	PinMode(s.a, Output)
	PinMode(s.b, Output)
	PinMode(s.c, Output)
	PinMode(s.d, Output)
	s.write(10)
}

// Writes the step value to the stepper
// The order:
// 10, 6, 5, 9
func (s Stepper) write(val int) {
	DigitalWrite(s.a, (val>>3)&1)
	DigitalWrite(s.b, (val>>2)&1)
	DigitalWrite(s.c, (val>>1)&1)
	DigitalWrite(s.d, val&1)
	s.step = val
}
