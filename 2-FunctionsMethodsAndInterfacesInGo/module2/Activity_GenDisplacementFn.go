package main

import (
	"fmt"
)

func main() {
	a := userInput("Enter acceleration: ")
	v0 := userInput("Enter initial velocity: ")
	s0 := userInput("Enter initial displacement: ")
	t := userInput("Enter time: ")

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement after %f time is %f", t, fn(t))
}

func GenDisplaceFn(acceleration float64, initVel float64, initDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return (0.5 * acceleration * (time * time)) + (initVel * time) + initDisplacement
	}
}

func userInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		fmt.Println(err)
	}
	return input
}
