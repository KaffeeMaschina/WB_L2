package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type EvictionAlgo interface {
	evict(c *Cache)
}

type Fifo struct{}

func (f *Fifo) evict(c *Cache) {
	fmt.Println("evict cache with First In, First Out strategy")
}

type Lru struct{}

func (l *Lru) evict(c *Cache) {
	fmt.Println("evict cache with Least Recently Used strategy")

}

type Lfu struct{}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("evict cache with Least Frequently Used strategy")
}

type Cache struct {
	storage      []byte
	capacity     int
	usedCapacity int
	evictionAlgo EvictionAlgo
}

func NewCache(capacity int, e EvictionAlgo) *Cache {
	return &Cache{storage: make([]byte, 0),
		capacity:     capacity,
		evictionAlgo: e}
}
func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}
func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.usedCapacity--
}
func (c *Cache) Add(data []byte) {
	if c.usedCapacity == c.capacity {
		c.evict()
	}
	c.usedCapacity++
	c.storage = append(c.storage, data...)
}
