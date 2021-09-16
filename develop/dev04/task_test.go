package main

import (
	"reflect"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		name  string
		word1 string
		word2 string
		want  bool
	}{
		{name: "anagram", word1: "ябеда", word2: "дабея", want: true},
		{name: "not anagram", word1: "ябеда", word2: "беда", want: false},
		{name: "empty", word1: "", word2: "", want: true},
		{name: "3 a", word1: "baraka", word2: "kabara", want: true},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAnagram(tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
func TestSortChars(t *testing.T) {
	testCases := []struct {
		name        string
		unsortedStr string
		want        string
	}{
		{name: "", unsortedStr: "fargo", want: "afgor"},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortChars(tt.unsortedStr); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSearch(t *testing.T) {
	testCases := []struct {
		name  string
		words *[]string
		want  *AnagramSet
	}{
		{"example", &[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			&AnagramSet{"листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}}},
		{"one word left", &[]string{"пятак", "тяпка", "слиток"},
			&AnagramSet{"пятак": {"пятак", "тяпка"}}},
		{"is low", &[]string{"пяТак", "ПяткА", "ТЯПкА", "Листок", "Слиток", "столик"},
			&AnagramSet{"листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}}},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
