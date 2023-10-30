package main

import (
	"fmt"
	"math"
	"strings"
)

type Shape struct {
	Type string
}

type Circle struct {
	Shape
	Radius float64
}

type Rectangle struct {
	Shape
	Width  float64
	Height float64
}

type Triangle struct {
	Shape
	Base   float64
	Height float64
}

func (c Circle) CalculateArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) CalculateArea() float64 {
	return r.Width * r.Height
}

func (t Triangle) CalculateArea() float64 {
	return 0.5 * t.Base * t.Height
}

func main() {
	fmt.Println("Welcome to the Shape Area Calculator!")

	for {
		fmt.Print("Enter the shape (circle/rectangle/triangle) or 'quit' to exit: ")
		var input string
		fmt.Scanln(&input)

		input = strings.ToLower(input)

		if input == "quit" {
			fmt.Println("Goodbye!")
			break
		} else if input == "circle" {
			fmt.Print("Enter the radius of the circle: ")
			var radius float64
			fmt.Scanln(&radius)
			circle := Circle{Shape: Shape{Type: "Circle"}, Radius: radius}
			area := circle.CalculateArea()
			fmt.Printf("Area of the circle: %.2f\n", area)
		} else if input == "rectangle" {
			fmt.Print("Enter the width of the rectangle: ")
			var width float64
			fmt.Scanln(&width)
			fmt.Print("Enter the height of the rectangle: ")
			var height float64
			fmt.Scanln(&height)
			rectangle := Rectangle{Shape: Shape{Type: "Rectangle"}, Width: width, Height: height}
			area := rectangle.CalculateArea()
			fmt.Printf("Area of the rectangle: %.2f\n", area)
		} else if input == "triangle" {
			fmt.Print("Enter the base of the triangle: ")
			var base float64
			fmt.Scanln(&base)
			fmt.Print("Enter the height of the triangle: ")
			var height float64
			fmt.Scanln(&height)
			triangle := Triangle{Shape: Shape{Type: "Triangle"}, Base: base, Height: height}
			area := triangle.CalculateArea()
			fmt.Printf("Area of the triangle: %.2f\n", area)
		} else {
			fmt.Println("Invalid shape. Please enter 'circle', 'rectangle', or 'triangle'.")
		}
	}
}
