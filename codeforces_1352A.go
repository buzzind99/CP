//go:build codeforces_1352A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	t := nextInt()
	num := make([]int, t)
	ans := make([][]int, t)
	for i := range t {
		n := nextInt()
		count := 0
		digit := 0.0
		var sol []int
		for n > 0 {
			mod := n%10
			if mod != 0 {
				sol = append(sol, mod*int(math.Pow(10.0, digit)))
				n -= mod
				count++
			} else {
				n /= 10
				digit++
			}
		}
		num[i], ans[i] = count, sol
	}

	for i := range t {
		fmt.Println(num[i])
		for j := range len(ans[i]) {
			fmt.Printf("%d ", ans[i][j])
			if j == len(ans[i])-1 {
				fmt.Println()
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1352/A
  Tags: implementation, math
  Rating: 800
*/
