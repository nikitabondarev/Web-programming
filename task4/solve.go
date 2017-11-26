package main

import (
"strings"
"unicode"
)


func RemoveEven(s []int) []int {
	result := make([]int, 0)
	for _, a := range s {
		if a%2 == 1 {
			result = append(result, a)
		}
	}
	return result
}


func PowerGenerator(a int) func() int {
	p := 1
	e := a
	gen := func() int {
		p *= e
		return p
	}
	return gen
}


func DifferentWordsCount(s string) int {
	word := make([]string, 0)
	m := make(map[string]int)
	for _, c := range s {
		if unicode.IsLetter(rune(c)) {
			word = append(word, string(unicode.ToLower(rune(c))))
		} else {
			if len(word) > 0 {
				m[strings.Join(word, "")]++
				word = make([]string, 0)
			}
		}
	}
	if len(word) > 0 {
		m[strings.Join(word, "")]++
	}
	return len(m)
}


