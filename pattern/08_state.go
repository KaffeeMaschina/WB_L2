package pattern

/*
	Реализовать паттерн «состояние».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/State_pattern
*/
type MobileAlertStater interface {
	Alert() string
}

type MobileAlert struct {
	state MobileAlertStater
}

func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

type MobileAlertVibration struct {
}

func (a *MobileAlertVibration) Alert() string {
	return "vibration"
}

type MobileAlertSong struct {
}

func (a *MobileAlertSong) Alert() string {
	return "Sing some song"
}
