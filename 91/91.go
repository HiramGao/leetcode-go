package main

func numDecodings(s string) int {
	n := len(s)
	//f[i-2] f[i-1] f[i]
	a, b, c := 0, 1, 0
	for i := 1; i <= n; i++ {
		c = 0
		if s[i-1] != '0' {
			c += b
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0')) <= 26 {
			c += a
		}
		a, b = b, c
	}

	return c
}
func main() {
	println(numDecodings("06"))
}
