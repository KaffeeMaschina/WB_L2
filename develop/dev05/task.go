package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func ParseFlagsAndArgs() (fl *Flags, pattern string, files []string) {
	flags := Flags{}
	flag.IntVar(&flags.after, "A", 0, "Print +N lines after each match")
	flag.IntVar(&flags.before, "B", 0, "Print +N lines before each match")
	flag.IntVar(&flags.context, "C", 0, "Print ±N lines before and after each match")
	flag.BoolVar(&flags.count, "c", false, "Print only a count of selected lines")
	flag.BoolVar(&flags.ignoreCase, "i", false, "Ignore case")
	flag.BoolVar(&flags.invert, "v", false, "Selected lines are those not matching any of the specified patterns")
	flag.BoolVar(&flags.fixed, "f", false, "Interpret patterns as fixed strings, not regular expressions")
	flag.BoolVar(&flags.lineNum, "n", false, "Print line number before each line")
	flag.Parse()
	pattern = flag.Arg(0)
	if pattern == "" {
		fmt.Println("Usage: grep [OPTIONS] PATTERN [FILE]")
		flag.PrintDefaults()
		os.Exit(1)
	}
	files = flag.Args()[1:]
	return &flags, pattern, files
}

func Grep(pattern string, fl *Flags, r io.Reader) string {
	var outputLines string
	scanner := bufio.NewScanner(r)
	var lines []string
	var linesCounter int
	for lineNumber := 1; scanner.Scan(); lineNumber++ {

		line := scanner.Text()

		if (Comparer(line, pattern, fl) && !fl.invert) || (!Comparer(line, pattern, fl) && fl.invert) {
			for i, prevLine := range lines {
				outputLines += outputLine(prevLine, lineNumber-(len(lines)-i), fl)
			}

			outputLines += outputLine(line, lineNumber, fl)

			for i := 1; i <= fl.after+fl.context; i++ {
				if scanner.Scan() {
					nextLine := scanner.Text()
					outputLines += outputLine(nextLine, lineNumber+i, fl)
				}
			}
			lines = nil
			linesCounter++
		} else if fl.before > 0 || fl.context > 0 {
			lines = append(lines, line)
			if len(lines) > fl.before+fl.context {
				lines = lines[1:]
			}
		}
	}
	if fl.count {
		outputLines += fmt.Sprintf("There are %v matches", linesCounter)
	}
	return outputLines
}
func Comparer(line, pattern string, fl *Flags) bool {
	if fl.fixed {
		if fl.ignoreCase {
			return strings.ToLower(line) == strings.ToLower(pattern)
		}
		return line == pattern
	}
	if fl.ignoreCase {
		return strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
	}
	return strings.Contains(line, pattern)
}
func outputLine(line string, lineNumber int, fl *Flags) string {
	if fl.lineNum {
		return fmt.Sprintf("%d:%s\n", lineNumber, line)
	} else {
		return fmt.Sprintln(line)
	}
}
func main() {
	fl, pattern, files := ParseFlagsAndArgs()
	if len(files) == 0 {
		r := io.Reader(os.Stdin)
		fmt.Println(Grep(pattern, fl, r))

	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Cannot open file: %v\n", err)
				continue
			}
			fmt.Println(Grep(pattern, fl, f))

		}
	}
}
