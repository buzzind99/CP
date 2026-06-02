//go:build codeforces_1660B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

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
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		a := make([]int, 0, n)
		for range n {
			a = append(a, nextInt())
		}

		if n == 1 {
			if a[0] == 1 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
			continue
		}

		slices.Sort(a)
		m1 := a[n-1]
		m2 := a[n-2]
		if m1 - m2 <= 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1660/B
  Tags: math
  Rating: 800
*/
