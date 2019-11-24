package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	fmt.Print("acceleration: ")
	stdin.Scan()
	_acceleration, _ := strconv.ParseFloat(stdin.Text(), 64)

	fmt.Print("initial velocity: ")
	stdin.Scan()
	_initialVelocity, _ := strconv.ParseFloat(stdin.Text(), 64)

	fmt.Print("initial displacement: ")
	stdin.Scan()
	_initialDisplacement, _ := strconv.ParseFloat(stdin.Text(), 64)

	fmt.Print("time: ")
	stdin.Scan()
	_time, _ := strconv.ParseFloat(stdin.Text(), 64)

	fn := GenDisplaceFn(_acceleration, _initialVelocity, _initialDisplacement)

	fmt.Println("displacement after", _time, "seconds:", fn(_time))
}
