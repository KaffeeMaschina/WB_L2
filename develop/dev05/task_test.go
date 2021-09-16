package main

import (
	"os"
	"testing"
)

const filename = "test.txt"

func TestGrep(t *testing.T) {
	testCases := []struct {
		name    string
		pattern string
		fl      Flags
		want    string
	}{
		{
			name: "test counter", fl: Flags{count: true}, pattern: "club",
			want: "I represent here a gopher club\nclub\nThe club is located in Saint-Petersburg\nThere are 3 matches"},
		{
			name: "test invertion", fl: Flags{invert: true}, pattern: "club",
			want: "Hello!\nMy name is Benjamin\nClub is not large\nand here goes mistake cLuB\n"},
		{
			name: "test ignore case", fl: Flags{ignoreCase: true}, pattern: "club",
			want: "I represent here a gopher club\nclub\nThe club is located in Saint-Petersburg\nClub is not large\nand here goes mistake cLuB\n"},
		{
			name: "test fixed string", fl: Flags{fixed: true}, pattern: "club",
			want: "club\n"},
		{
			name: "test line numbers", fl: Flags{lineNum: true}, pattern: "club",
			want: "3:I represent here a gopher club\n4:club\n5:The club is located in Saint-Petersburg\n"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(filename)
			if err != nil {
				t.Fatal(err)
			}
			got := Grep(tt.pattern, &tt.fl, f)
			if got != tt.want {
				t.Errorf("got '%v', want '%v'", got, tt.want)
			}
		})
	}
}
func TestComparer(t *testing.T) {
	testCases := []struct {
		name    string
		line    string
		pattern string
		fl      Flags
		want    bool
	}{
		{name: "no flags test", line: "My name is Benjamin", pattern: "name", want: true},
		{name: "test fixed line", line: "My name is Benjamin", pattern: "name", fl: Flags{fixed: true}, want: false},
		{name: "test ignore case", line: "BOYZ ClUb", pattern: "club", fl: Flags{ignoreCase: true}, want: true},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Comparer(tt.line, tt.pattern, &tt.fl)
			if got != tt.want {
				t.Errorf("got '%v', want '%v'", got, tt.want)
			}
		})
	}
}
