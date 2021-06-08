package main

import "fmt"

func removeInvalidParentheses(s string) []string {
	n := len(s)
	leftRemove, rightRemove := 0, 0
	hashMap := make(map[string]bool)

	var dfs func(index, leftCount, RightCount, leftRemove, rightRemove int, path []byte)

	/**
	 * @param index       当前遍历到的下标
	 * @param leftCount   已经遍历到的左括号的个数
	 * @param rightCount  已经遍历到的右括号的个数
	 * @param leftRemove  最少应该删除的左括号的个数
	 * @param rightRemove 最少应该删除的右括号的个数
	 * @param path        一个可能的结果
	 */

	dfs = func(index, leftCount, rightCount, leftRemove, rightRemove int, path []byte) {
		if index == n {
			if leftRemove == 0 && rightRemove == 0 {
				hashMap[string(path)] = true
			}
			return
		}
		currentByte := s[index]
		//1 删除当前字符串
		if currentByte == '(' && leftRemove > 0 {
			dfs(index+1, leftCount, rightCount, leftRemove-1, rightRemove, path)
		}
		if currentByte == ')' && rightRemove > 0 {
			dfs(index+1, leftCount, rightCount, leftRemove, rightRemove-1, path)
		}
		//2 保留当前的字符
		path = append(path, currentByte)
		if currentByte != '(' && currentByte != ')' {
			dfs(index+1, leftCount, rightCount, leftRemove, rightRemove, path)
		} else if currentByte == '(' {
			dfs(index+1, leftCount+1, rightCount, leftRemove, rightRemove, path)
		} else if rightCount < leftCount {
			dfs(index+1, leftCount, rightCount+1, leftRemove, rightRemove, path)
		}
		path = path[:len(path)-1]
	}

	for _, b := range []byte(s) {
		if b == '(' {
			leftRemove++
		} else if b == ')' {
			if leftRemove == 0 {
				rightRemove++
			}
			if leftRemove > 0 {
				leftRemove--
			}
		}
	}

	dfs(0, 0, 0, leftRemove, rightRemove, []byte{})
	res := []string{}
	for key := range hashMap {
		res = append(res, key)
	}

	return res
}

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("(a)())()"))
	fmt.Println(removeInvalidParentheses(")("))
	fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("(a)())()"))
}
