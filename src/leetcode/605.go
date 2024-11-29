// Can place flowers
package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++{
		// In case task is complete before reaching the end of the array
		if n == 0 {
			return true
		}
		if flowerbed[i] == 0 {
			// Check if the current spot is valid to plant a flower
			prevEmpty := i == 0 || flowerbed[i - 1] == 0
			nextEmpty := i == len(flowerbed) - 1 || flowerbed[i + 1] == 0
			if prevEmpty && nextEmpty {
				flowerbed[i] = 1 // Plant a flower
				n-- // Reduce the number of flowers to plant
			}
		}
	}
	return n == 0
}