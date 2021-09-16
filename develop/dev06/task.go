package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

# Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
const (
	NoFieldError         = "A number of field should be mentioned"
	WrongNumbersOfFields = "Option field should contain numbers of field separated with ','"
)

type Flags struct {
	fields    []int
	delimiter string
	separated bool
}

func ParseFlagsAndArgs() (flags *Flags) {
	fl := Flags{}
	var fields string
	flag.StringVar(&fields, "f", "0", "choose field")
	flag.StringVar(&fl.delimiter, "d", "", "field delimiter")
	flag.BoolVar(&fl.separated, "s", false, "fields with delimiter only")
	flag.Parse()
	if fields == "0" {
		flag.PrintDefaults()
		log.Fatal(errors.New(NoFieldError))
	}
	for _, field := range strings.Split(fields, ",") {
		numField, err := strconv.Atoi(field)
		if err != nil {
			log.Fatal(errors.New(WrongNumbersOfFields))
		}
		fl.fields = append(fl.fields, numField)
	}

	return &fl
}

func Cut(fl *Flags, r io.Reader) []string {
	var output []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		var outLine string
		if fl.separated && strings.Contains(line, fl.delimiter) {
			outLine = SelectData(line, fl.fields, fl.delimiter)
			output = append(output, outLine)

		} else {
			outLine = SelectData(line, fl.fields, fl.delimiter)
			output = append(output, outLine)
		}

	}
	return output
}
func SelectData(line string, fields []int, delim string) string {
	if delim == "" {
		delim = "\t"
	}
	subStrings := strings.Split(line, delim)
	var outputLines []string
	for _, numField := range fields {

		if len(subStrings) >= numField {
			outputLines = append(outputLines, subStrings[numField-1])
		}
	}
	return strings.Join(outputLines, delim)

}
func Output(lineToPrint string) {
	fmt.Println(lineToPrint)
}
func main() {
	fl := ParseFlagsAndArgs()

	output := Cut(fl, os.Stdin)
	fmt.Println(output)
	for _, line := range output {
		Output(line)
	}
}
