package go_basic

import "strconv"

// 只出现一次的数字
func SingleNumber(nums []int) int {
	checkMap := make(map[int]int)
	for _, num := range nums {
		checkMap[num]++
	}

	for num, count := range checkMap {
		if count == 1 {
			return num
		}
	}
	return 0
}

// 回文数
func IsPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	length := len(strX)
	for i := 0; i < length/2; i++ {
		if strX[i] != strX[length-i-1] {
			return false
		}
	}
	return true
}
