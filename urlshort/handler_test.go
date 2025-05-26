package urlshort

import (
	"fmt"
	"log"
	"testing"
)

func TestParseYAML(t *testing.T) {
	yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	got, err := parseYAML([]byte(yml))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(got)
}
