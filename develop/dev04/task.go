package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
type AnagramSet map[string]*[]string

func main() {
	arr := []string{"ябеда", "беда", "бедая", "даябе", "дабе"}
	arr2 := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	fmt.Println(Search(&arr))
	fmt.Println(Search(&arr2))
}

func Search(words *[]string) *AnagramSet {
	res := make(AnagramSet)
	for _, word := range *words {
		isStored := false
		word = strings.ToLower(word)
		//sortedWord := SortChars(word)

		for k, _ := range res {
			if IsAnagram(word, k) {
				isStored = true
				*res[k] = append(*res[k], word)
				break
			}
		}
		if !isStored {
			res[word] = &[]string{word}
		}

	}
	for k, v := range res {
		if len(*v) == 1 {
			delete(res, k)
		}
	}
	return &res

}
func SortChars(unsortedStr string) string {
	chars := strings.Split(unsortedStr, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
func IsAnagram(first, second string) bool {
	return SortChars(first) == SortChars(second)
}
