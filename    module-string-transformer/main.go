package main

import (
	"fmt"
	"strings"
)
func capEachWord(word string) string {
	return strings.ToUpper(word) 
}

func lower(word string) string {
	return strings.ToLower(word)
}

func cap(word string) string {
	return strings.Title(strings.ToLower(word))
}

func TitleCase(s string) string {
	w := strings.Fields(S)

	for i, j := range w {
		if len(j) > 3 {
			w[i] = strings.ToUpper(j)[:1] + strings.ToLower(j[1:])
		}
	}
}

func Reverse(word string) string {
	s := strings.Fields(word)
	
	for i, j := range s {
		x := []rune (j)
		for a,b := 0, len(x)-1; a < b; a,b = a+1, b-1 {
		x[a], x[b] = x[b], x[a]
	}
	s[i] = string(x)
}
    output := strings.Join(s, " ")
	return output
}
func main() {

	fmt.Println(capEachWord("sentinel is online"))
	fmt.Println(lower("ALERT LEVEL FIVE DETECTED"))
	fmt.Println(cap("director adaeze okonkwo"))
	fmt.Println(cap("THREAT LEVEL elevated"))
	fmt.Println(Reverse("Lagos Nigeria"))
}