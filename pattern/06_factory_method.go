package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
type Weapon interface {
	Attack()
}

type Spear struct {
	HitPoints int
	Amount    int
}

func (s *Spear) Attack() {
	fmt.Printf("spear attack with power %d\n", s.HitPoints+s.Amount)
}

type Sword struct {
	HitPoints int
	Amount    int
}

func (sw *Sword) Attack() {
	fmt.Printf("sword attack with power %d\n", sw.HitPoints+sw.Amount)
}

type WeaponFactory interface {
	CreateWeapon(amount int)
}
type SpearFactory struct{}

func (sf *SpearFactory) CreateWeapon(amount int) Weapon {
	fmt.Printf("Create %d spears\n ", amount)
	return &Spear{HitPoints: 5,
		Amount: amount}
}

type SwordFactory struct{}

func (swf *SwordFactory) CreateWeapon(amount int) Weapon {
	fmt.Printf("Create %d swords\n ", amount)
	return &Sword{HitPoints: 10,
		Amount: amount}
}
