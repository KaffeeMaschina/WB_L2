Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Программа выведет:
2
1
В случае test(), defer, выполняясь после return, модифицирует значение x, которое объявленно
в сигнатуре функции.
Во случае anotherTest(), defer, модицфицирует локальную переменную x. В то время как return 
уже вернул значение.
Это можно проверить если добавить ещё одну отложенную функцию, которая будет печатать значение 
переменной x после всех изменений:

func anotherTest() int {
	var x int
	defer func() {
		fmt.Println(x)
	}()
	defer func() {
		x++
	}()
	x = 1
	return x
}

fmt.Println(anotherTest()) 
//output:
2   // defer, который печатает x
1   // результат функции


```
