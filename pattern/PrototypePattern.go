package pattern

import "fmt"

// ShapeType Custom type defined
type ShapeType int

const (
	CircleType ShapeType = 1
	SquareType ShapeType = 2
)

type Shape interface {
	GetId() ShapeType // Get the ShapeId
	PrintTypeProp()   // Print property values of the Shape
	Clone() Shape     // Getting Deep Copy of the Shape
}

// Circle Implementation of Shape
type Circle struct {
	Id            ShapeType
	Radius        int
	Diameter      int
	Circumference int
}

func NewCircle(radius, diameter, circumference int) Circle {
	return Circle{CircleType, radius, diameter, circumference}
}

func (c Circle) GetId() ShapeType {
	return c.Id
}

func (c Circle) Clone() Shape { // Prototyping
	return NewCircle(c.Radius, c.Diameter, c.Circumference)
}

func (c Circle) PrintTypeProp() {
	fmt.Println("Circle Properties Radius:", c.Radius, "Diameter:", c.Diameter, "Circumference:", c.Circumference)
}

type Square struct {
	Id      ShapeType
	Length  int
	Breadth int
}

func NewSquare(length, breadth int) Square {
	return Square{SquareType, length, breadth}
}

func (s Square) GetId() ShapeType {
	return s.Id
}

func (s Square) Clone() Shape { // Prototyping
	return NewSquare(s.Length, s.Breadth)
}

func (s Square) PrintTypeProp() {
	fmt.Println("Square Properties Length:", s.Length, "Breadth:", s.Breadth)
}

var RegistryList = make(map[int]Shape) // Prototype Registry

func LoadToRegistry() {
	circle := NewCircle(10, 20, 30)
	RegistryList[int(circle.GetId())] = circle // Registering Circle to Registry

	square := NewSquare(10, 20)
	RegistryList[int(square.GetId())] = square // Registering Square to Registry
}

func GetPrototypeName(s string) string {
	return "Code for : -- " + s + " Pattern --"
}
