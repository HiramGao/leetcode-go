package main

func minArray(numbers []int) int {
	i, j := 0, len(numbers)-1
	mid := i
	for numbers[i] >= numbers[j] {
		if j-i == 1 {
			return numbers[j]
		}
		mid = (i + j) / 2
		if numbers[i] == numbers[mid] && numbers[mid] == numbers[j] {
			res := numbers[i]
			for ; i <= j; i++ {
				if res > numbers[i] {
					res = numbers[i]
				}
			}
			return res
		}
		if numbers[mid] >= numbers[i] {
			i = mid
		} else if numbers[mid] <= numbers[j] {
			j = mid
		}
	}
	return numbers[mid]
}

func main() {
	println(minArray([]int{2, 2, 2, 0, 1}))
}
