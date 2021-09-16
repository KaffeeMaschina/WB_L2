package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	testCases := []struct {
		name   string
		fl     *Flags
		reader io.Reader
		want   []string
	}{
		{
			name:   "second field, default delimiter, not separated",
			fl:     &Flags{fields: []int{2}},
			reader: strings.NewReader("hello\tmy\tname\nis\tNikita\tAnd\nI'm\tstudy\tGolang\n1\t2\t3\n"),
			want:   []string{"my", "Nikita", "study", "2"},
		},
		{
			name: "second and third field, default delimiter, not separated",
			fl:   &Flags{fields: []int{2, 3}},
			reader: strings.NewReader("hello\tmy\tname\n" +
				"is\tNikita\tAnd\n" +
				"I'm\tstudy\tGolang\n" +
				"1\t2\t3\n"),
			want: []string{"my\tname",
				"Nikita\tAnd",
				"study\tGolang",
				"2\t3"},
		},
		{
			name: "second and third field, ',' delimiter, not separated",
			fl:   &Flags{fields: []int{2, 3}, delimiter: ","},
			reader: strings.NewReader("hello,my,name\n" +
				"is,Nikita,And\n" +
				"I'm,study,Golang\n" +
				"1,2,3\n"),
			want: []string{"my,name",
				"Nikita,And",
				"study,Golang",
				"2,3"},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Cut(tt.fl, tt.reader)
			want := tt.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}

		})
	}
}
func TestSelectData(t *testing.T) {
	testCases := []struct {
		name   string
		line   string
		fields []int
		delim  string
		want   string
	}{
		{
			name:   "field 4 of 3",
			line:   "hello\tmy\tname\n",
			fields: []int{4},
			delim:  "",
			want:   "",
		},
		{
			name:   "delim '-', ignore '\t'",
			line:   "hello\tmy-name\n",
			fields: []int{2},
			delim:  "-",
			want:   "name\n",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectData(tt.line, tt.fields, tt.delim)
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
