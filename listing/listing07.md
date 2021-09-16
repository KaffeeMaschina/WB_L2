Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Поеведение будет не предсказуемым.
Сначала будут выведены числа которые мы передаём, но возможно не все, и порядок вывода каждой последовательности не будет нарушаться.
Однако числа одной последовательноти могут выводиться быстрее.
В блоке select выбирается case с той операцией которая является не блокирующей, если оба канала будут готовы отдать значение,
то выбор будет произвольным.
После того как все числа будут напечатаны, канал закроется и из него будут бесконечно приходить нулевые значения.
Мы может защититься от нулевых значений, если будем использовать bool переменную, которая будет передавать значение о закрытости канала:
case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
И цикл будет работать только есть хоть один из каналов не равен nil.
Так же нужно закрыть канал, что бы не поймать deadlock.
go func() {
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
```

