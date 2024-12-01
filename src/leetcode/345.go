// reverse vowels of a string
package main

func isVowel(c string) bool {
  vowels := map[string]bool {
    "a": true, "e": true, "i": true, "o":true, "u": true,
    "A": true, "E": true, "I": true, "O":true, "U": true,
  }
  return vowels[c]
}

func reverseVowels(s string) string {
  left := 0
  right := len(s) - 1
  leftPart := ""
  rightPart := ""
  for left < right {
    for left < right && !isVowel(string(s[left])) {
      leftPart += string(s[left])
      left++
    }
    for left < right && !isVowel(string(s[right])) {
      rightPart = string(s[right]) + rightPart
      right--
    }
    if left < right && isVowel(string(s[left])) && isVowel(string(s[right])){
      leftPart += string(s[right])
      rightPart = string(s[left]) + rightPart
      left++
      right--
    }
  }
  if len(rightPart + leftPart) < len(s) {
    return leftPart + string(s[left]) + rightPart
  } else {
    return leftPart + rightPart
  }
}

func isVowelByte(c byte) bool {
  vowels := map[byte]bool {
    'a': true, 'e': true, 'i': true, 'o':true, 'u': true,
    'A': true, 'E': true, 'I': true, 'O':true, 'U': true,
  }
  return vowels[c]
}

func reverseVowelsByte(s string) string {
  bytes := []byte(s)
  left := 0
  right := len(s) - 1
  for left < right {
    for left < right && !isVowelByte(bytes[left]) {
      left++
    }
    for left < right && !isVowelByte(bytes[right]) {
      right--
    }
    temp:= bytes[left]
    bytes[left] = bytes[right]
    bytes[right] = temp
    left++
    right--
  }
  return string(bytes)
}