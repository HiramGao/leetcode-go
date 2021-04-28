package main

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1*l2 == 0 {
		return l1 + l2
	}

	f := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		f[i] = make([]int, l2+1)
	}
	for i := 0; i <= l1; i++ {
		f[i][0] = i
	}
	for i := 0; i <= l2; i++ {
		f[0][i] = i
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i-1][j-1], f[i-1][j], f[i][j-1]) + 1
			}
		}
	}
	return f[l1][l2]
}

func min(arr ...int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min
}
func main() {
	println(minDistance("distance", "springbok"))
}
