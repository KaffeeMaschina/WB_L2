выведетЧто выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

Структура customError подходит под интерфейс error, потому что реализует метод Error() string.
Поэтому мы можем присвоить результат функции test() переменной err.
Функция test() в свою очередь возвращает указатель, который уже проинициализирован, и соотвественно
не может быть nil.
Если мы заменим в функции test() *customError на error, в таком случае, результат будет nil
```






