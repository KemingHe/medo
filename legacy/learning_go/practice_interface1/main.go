package main

func main() {
	trianglePtr := &triangle{
		height: 1,
		base: 2,
	}
	
	squarePtr := &square{
		sideLength: 1,
	}

	printShapeArea(trianglePtr)
	printShapeArea(squarePtr)
}
