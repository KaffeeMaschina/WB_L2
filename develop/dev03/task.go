package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	filename string
	column   int
	number   bool
	revers   bool
	uniq     bool
}

func ParseFlags() *Flags {
	f := Flags{}

	flag.StringVar(&f.filename, "f", "", "filename to sort")
	flag.IntVar(&f.column, "k", -1, "column to sort")
	flag.BoolVar(&f.number, "n", false, "sort by number")
	flag.BoolVar(&f.revers, "r", false, "reverse sort")
	flag.BoolVar(&f.uniq, "u", false, "sort unique strings")
	flag.Parse()

	return &f
}

func readFile(filename string) []string {
	var lines []string
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
func SortByOptions(lines []string, fl *Flags) []string {

	comparator := func(i, j int) bool {
		a, b := lines[i], lines[j]

		if fl.column > 0 {
			aCols := strings.Fields(a)
			bCols := strings.Fields(b)

			if fl.column <= len(aCols) && fl.column <= len(bCols) {
				a = aCols[fl.column-1]
				b = bCols[fl.column-1]
			}
		}
		if fl.number {
			aNum, errA := strconv.ParseFloat(a, 64)
			bNum, errB := strconv.ParseFloat(b, 64)
			if errA == nil && errB == nil {
				return aNum < bNum
			}
		}
		return a < b
	}
	if fl.uniq {
		lines = removeDuplicates(lines)
	}
	if fl.revers {
		sort.SliceStable(lines, func(i, j int) bool {
			return !comparator(i, j)
		})
	} else {
		sort.SliceStable(lines, comparator)
	}

	return lines
}
func removeDuplicates(lines []string) []string {
	uniqLines := make(map[string]struct{})
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		if _, ok := uniqLines[line]; !ok {
			uniqLines[line] = struct{}{}
			result = append(result, line)
		}
	}
	return result
}

func output(data []string) {
	for i := range data {
		fmt.Println(data[i])
	}
}
func main() {
	fl := ParseFlags()
	lines := readFile(fl.filename)
	res := SortByOptions(lines, fl)
	output(res)
}
