package main

//func isMatch(s string, p string) bool {
//	n, m := len(s), len(p)
//	f := make([][]bool, n+1)
//	for i := range f {
//		f[i] = make([]bool, m+1)
//	}
//
//	for i := 0; i <= n; i++ {
//		for j := 0; j <= m; j++ {
//			if j == 0 {
//				f[i][0] = i == 0
//			} else {
//				if p[j-1] != '*' {
//					if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
//						f[i][j] = f[i-1][j-1]
//					}
//				} else {
//					if j >= 2 {
//						f[i][j] = f[i][j] || f[i][j-2]
//					}
//					if i >= 1 && j >= 2 && (s[i-1] == p[j-2] || p[j-2] == '.') {
//						f[i][j] = f[i][j] || f[i-1][j]
//					}
//				}
//			}
//		}
//	}
//
//	return f[n][m]
//}

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	//首字母是否相同
	res := !(len(s) == 0) && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (res && isMatch(s[1:], p))
	} else {
		return res && isMatch(s[1:], p[1:])
	}
}

func main() {
	println(isMatch("ab", "ab*"))
}
