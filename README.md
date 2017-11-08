# Go wiringPi
[![GoDoc](https://godoc.org/github.com/ilango100/go-wPi?status.svg)](https://godoc.org/github.com/ilango100/go-wPi)

go-wPi is a simple wrapper for accessing wiringPi functions via golang.

## Example

```
package main

import(
    "github.com/ilango100/go-wPi"
)

func main() {
    wPi.wiringPiSetup()
    wPi.PinMode(2,wPi.Output)
    wPi.DigitalWrite(2,1)
}
```

You can also add stepper type to easily control a stepper motor.
```
motor := wPi.NewStepper(200, 5, 6, 13, 19)
motor.Step(20)
```

## Bugs / Contributing
Feel free to fork and create pull request and also to open an issue if you face one.
