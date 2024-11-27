// Greated common divisor of strings
package main

func gcdOfStrings(str1 string, str2 string) string {
	if str1 + str2 != str2 + str1 {
		return ""
	} else {
		a := len(str1)
		b := len(str2)
		var temp int
		for b > 0 {
			temp = a % b
			a = b
			b = temp
		}
		return str1[0:a]
	}
}