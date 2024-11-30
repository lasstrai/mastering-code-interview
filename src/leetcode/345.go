// reverse vowels of a string
package main

import (
	"strings"
)

func reverseVowels(s string) string {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	left := 0
	right := len(s) - 1
  result := ""
	for left <= right {
		currentLeft := string(s[left])
    currentRight := string(s[right])
		_, existsLeft := vowels[strings.ToLower(currentLeft)]
    _, existsRight := vowels[strings.ToLower(currentRight)]
    if existsLeft && existsRight {
      result = string(s[right]) + result
      result = result + string(s[left])
      left++
      right--
    }
    if !existsLeft {
      result += string(s[left])
      left++
    }
    if !existsRight {
      result += string(s[right])
      right--
    }
	}
  return result
}