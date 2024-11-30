package main

func removeElement(nums []int, val int) int {
	i := -1
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			i++
			if i < j {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	return i + 1
}