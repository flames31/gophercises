package parse

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestParseHTML(t *testing.T) {
	filepath := "/Users/rahul31/Desktop/gophercises/link/html_files"
	type testCase struct {
		input *os.File
		want  []Link
	}
	cases := []testCase{}

	for i := 1; i <= 2; i++ {
		fileName := fmt.Sprintf("%v/ex%v.html", filepath, i)
		answers := getAnswers()
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("ERR : %v", err)
		}
		defer file.Close()
		cases = append(cases, testCase{
			input: file,
			want:  answers[i],
		})
	}

	for _, c := range cases {
		got, _ := ParseHTML(c.input)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf(`---------------------------------
	input:  %v
	want:   %v
	got:    %v
	`, c.input, c.want, got)
		}
	}
}
