package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

// This is wrong: func (s shape) printArea...

func printShapeArea(s shape) {
	area := s.getArea()

	// Use type switch to determine the type name
	var typeName string
	switch s.(type) {
	case *square:
		typeName = "square"
	case *triangle:
		typeName = "triangle"
	default:
		typeName = "shape"
	}

	fmt.Printf("Area of %s is %v\n", typeName, area)
}

func (ptr *triangle) getArea() float64 {
	return (*ptr).base * (*ptr).height * 0.5
}

func (ptr *square) getArea() float64 {
	return math.Pow((*ptr).sideLength, 2)
}
