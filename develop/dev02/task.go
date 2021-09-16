package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var ErrWrongString = errors.New("incorrect string")

func main() {
	fmt.Println(unpack(`4asd4bk2`))

}
func unpack(s string) (string, error) {

	var sb strings.Builder
	for i, v := range s {
		if unicode.IsDigit(rune(s[0])) || unicode.IsDigit(v) && unicode.IsDigit(rune(s[i-1])) {
			return "", ErrWrongString
		}
		if unicode.IsDigit(v) {
			num, err := strconv.ParseInt(string(v), 10, 0)
			if err != nil {
				return "", fmt.Errorf("cannot ParseInt :%w", err)
			}
			for j := 0; j < int(num)-1; j++ {
				sb.WriteByte(s[i-1])
			}
		} else {
			sb.WriteByte(s[i])
		}

	}
	return sb.String(), nil
}
