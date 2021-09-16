package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Visitor_pattern
*/
type Building interface {
	GetType() string
	Accept(Visitor)
}
type Visitor interface {
	visitForChurch(*Church)
	visitForHotel(*Hotel)
	visitForStadium(*Stadium)
}
type Church struct {
	sideA int
	sideB int
}

func (c *Church) Accept(v Visitor) {
	v.visitForChurch(c)
}
func (c *Church) GetType() string {
	return "Church"
}

type Hotel struct {
	l int
	b int
}

func (h *Hotel) Accept(v Visitor) {
	v.visitForHotel(h)
}
func (h *Hotel) GetType() string {
	return "Hotel"
}

type Stadium struct {
	radius int
}

func (s *Stadium) Accept(v Visitor) {
	v.visitForStadium(s)
}
func (s *Stadium) GetType() string {
	return "Stadium"
}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForChurch(c *Church) {
	fmt.Println("calculating area for church")
}
func (a *AreaCalculator) visitForHotel(h *Hotel) {
	fmt.Println("calculating area for hotel")
}
func (a *AreaCalculator) visitForStadium(s *Stadium) {
	fmt.Println("calculating area for stadium")
}

type MiddleCoordinates struct {
	x int
	y int
}

func (m *MiddleCoordinates) visitForChurch(c *Church) {
	fmt.Println("calculating middle point for church")
}
func (m *MiddleCoordinates) visitForHotel(h *Hotel) {
	fmt.Println("calculating middle point for hotel")
}
func (m *MiddleCoordinates) visitForStadium(s *Stadium) {
	fmt.Println("calculating middle point for stadium")
}

func Visit() {
	church := &Church{sideA: 1, sideB: 2}
	hotel := &Hotel{l: 2, b: 3}
	stadium := &Stadium{radius: 5}

	areaCalculator := &AreaCalculator{}
	church.Accept(areaCalculator)
	hotel.Accept(areaCalculator)
	stadium.Accept(areaCalculator)

	middleCoordinates := &MiddleCoordinates{}
	church.Accept(middleCoordinates)
	hotel.Accept(middleCoordinates)
	stadium.Accept(middleCoordinates)
}
