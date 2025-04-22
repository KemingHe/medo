package main

import (
	"io"
	"os"
	"testing"
)

const testTriangle1Height = 2
const testTriangle1Base = 5
const testTriangle1Area = 5
const testTriangle1ExpectedOutput = "Area of triangle is 5\n"

const testSquare1SideLength = 3
const testSquare1Area = 9
const testSquare1ExpectedOutput = "Area of square is 9\n"

func TestTriangle(t *testing.T) {
	tPtr := &triangle{
		height: testTriangle1Height,
		base:   testTriangle1Base,
	}

	tHeight := (*tPtr).height
	if tHeight != testTriangle1Height {
		t.Errorf("Expected height initialized as %v, got %v", testTriangle1Height, tHeight)
	}

	tBase := (*tPtr).base
	if tBase != testTriangle1Base {
		t.Errorf("Expected base initialized as %v, got %v", testTriangle1Base, tBase)
	}

	tArea := tPtr.getArea()
	if tArea != testTriangle1Area {
		t.Errorf("Expected triangle area of %v, got %v", testTriangle1Area, tArea)
	}
}

func TestSquare(t *testing.T) {
	tPtr := &square{
		sideLength: testSquare1SideLength,
	}

	tSideLength := (*tPtr).sideLength
	if tSideLength != testSquare1SideLength {
		t.Errorf("Expected side length initialized as %v, got %v\n", testSquare1SideLength, tSideLength)
	}

	tArea := tPtr.getArea()
	if tArea != testSquare1Area {
		t.Errorf("Expected square area of %v, got %v\n", testSquare1Area, tArea)
	}
}

func TestPrintShapeArea(t *testing.T) {
	// Create test shapes
	square1 := &square{sideLength: testSquare1SideLength}
	triangle1 := &triangle{height: testTriangle1Height, base: testTriangle1Base}

	// Test each shape
	testShapePrinting(t, square1, testSquare1ExpectedOutput)
	testShapePrinting(t, triangle1, testTriangle1ExpectedOutput)
}

// Helper function to test printShapeArea for any shape
func testShapePrinting(t *testing.T, s shape, expected string) {
	// Redirect stdout to capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function we're testing
	printShapeArea(s)

	// Close the writer to get all output
	w.Close()

	// Read captured output
	var output []byte
	output, _ = io.ReadAll(r)

	// Restore stdout
	os.Stdout = oldStdout

	// Compare actual vs expected output
	// use %q instead of %v to explicitly show whitespace in string
	actual := string(output)
	if actual != expected {
		t.Errorf("For shape %T: expected output %q, got %q", s, expected, actual)
	}
}
