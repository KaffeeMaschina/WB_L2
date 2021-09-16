package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Command_pattern
*/
type Command interface {
	Execute()
}
type LightReceiver struct{}

func (l *LightReceiver) TurnOn() {
	fmt.Println("Light is on")
}
func (l *LightReceiver) TurnOff() {
	fmt.Println("Light is off")
}

type LightOnCommand struct {
	light *LightReceiver
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

type LightOffCommand struct {
	light *LightReceiver
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}
func (r *RemoteControl) PressButton() {
	r.command.Execute()
}
func client() {
	light := &LightReceiver{}
	onCommand := &LightOnCommand{light: light}
	offCommand := &LightOffCommand{light: light}
	remote := &RemoteControl{}
	remote.SetCommand(onCommand)
	remote.PressButton()
	remote.SetCommand(offCommand)
	remote.PressButton()
}
