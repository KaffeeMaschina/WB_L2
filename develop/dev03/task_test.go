package main

import (
	"reflect"
	"testing"
)

func TestSortByOptions(t *testing.T) {
	testCases := []struct {
		name string
		fl   *Flags
		want []string
	}{
		{
			name: "test-k",
			fl:   &Flags{filename: "test-k.txt", column: 2},
			want: []string{"Thomas 0 pear", "Luther 1 pineapple", "Martin 4 orange", "Martin 9 apple"},
		},
		{
			name: "test-n",
			fl:   &Flags{filename: "test-n.txt", number: true},
			want: []string{"0", "5", "10", "11", "12", "22"},
		},
		{
			name: "test-r",
			fl:   &Flags{filename: "test-r.txt", revers: true},
			want: []string{"Thomas", "Robert", "Patrick", "Martin", "John", "Anton"},
		},
		{
			name: "test-u",
			fl:   &Flags{filename: "test-u.txt", uniq: true},
			want: []string{"Martin", "Robert", "Thomas"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fl.column > 0 {
				got := SortByOptions(readFile(tt.fl.filename), tt.fl)
				want := tt.want
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			}
		})
	}
}
