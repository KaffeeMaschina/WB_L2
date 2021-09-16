package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».

Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Heater struct {
}

func (h *Heater) TurnOnHeater() {
	fmt.Println("Heating...")
}

type AirVent struct {
}

func (a *AirVent) TurnOnAirVentilation() {
	fmt.Println("Ventilating")
}

type Light struct {
}

func (l *Light) TurnOnLight() {
	fmt.Println("Light")
}

type OvenFacade struct {
	heater  *Heater
	airVent *AirVent
	light   *Light
}

func NewOvenFacade() *OvenFacade {
	return &OvenFacade{
		heater:  &Heater{},
		airVent: &AirVent{},
		light:   &Light{},
	}
}
func (o *OvenFacade) Bake() {
	o.heater.TurnOnHeater()
	o.airVent.TurnOnAirVentilation()
	o.light.TurnOnLight()
}
