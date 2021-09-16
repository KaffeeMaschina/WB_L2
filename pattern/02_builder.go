package pattern

/*
	Реализовать паттерн «строитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Builder_pattern
*/
type Car struct {
	color         string
	engine        string
	hasNavigation bool
}
type CarBuilder interface {
	SetColor(color string) CarBuilder
	SetEngine(engine string) CarBuilder
	SetNavigation(hasNavigation bool) CarBuilder
	Build() *Car
}
type carBuilder struct {
	car *Car
}

func (cb *carBuilder) SetColor(color string) CarBuilder {
	cb.car.color = color
	return cb
}
func (cb *carBuilder) SetEngine(engine string) CarBuilder {
	cb.car.engine = engine
	return cb
}
func (cb *carBuilder) SetNavigation(hasNavigation bool) CarBuilder {
	cb.car.hasNavigation = hasNavigation
	return cb
}
func (cb *carBuilder) Build() *Car {
	return cb.car
}
func NewCarBuilder() CarBuilder {
	return &carBuilder{
		car: &Car{},
	}
}

type Director struct {
	builder *carBuilder
}

func NewDirector(builder *carBuilder) *Director {
	return &Director{builder: builder}
}
func (d *Director) ConstructCar(color, engine string, hasNavigation bool) *Car {
	d.builder.SetColor(color).
		SetEngine(engine).
		SetNavigation(hasNavigation)

	return d.builder.Build()
}
