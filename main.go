package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
func main() {
	var s string
	fmt.Scanln(&s)
	tmp := strings.Split(s, ",")
	s1, s2 := tmp[0], tmp[1]
	a, _ := strconv.Atoi(tmp[2])
	b, _ := strconv.Atoi(tmp[3])
	c, _ := strconv.Atoi(tmp[4])
	m, n := len(s1), len(s2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = i * b
	}
	for j := 0; j < n; j++ {
		dp[0][j] = j * a
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			min := Min(dp[i-1][j]+b, dp[i][j-1]+a)
			if s1[i] == s2[j] {
				dp[i][j] = Min(min, dp[i-1][j-1])
			} else {
				dp[i][j] = Min(min, dp[i-1][j-1]+c)
			}
		}
	}
	fmt.Println(dp[m-1][n-1])
}
