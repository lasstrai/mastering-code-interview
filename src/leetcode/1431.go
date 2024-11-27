// Kids With the Greatest Number of Candies
package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, value := range candies {
		if value > max {
			max = value
		}
	}
	result := make([]bool, len(candies))
	for i, value := range(candies) {
		if value + extraCandies >= max {
			result[i] = true
		}
	}
	return result
}