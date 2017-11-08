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

func (s Stepper) init() {
	s.step = 10
	PinMode(s.a, Output)
	PinMode(s.b, Output)
	PinMode(s.c, Output)
	PinMode(s.d, Output)
}
