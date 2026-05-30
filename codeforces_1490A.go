//go:build codeforces_1490A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		ans := 0
		for i := 0; i < n-1; i++ {
			cur := min(a[i], a[i+1])
			target := max(a[i], a[i+1])
			
			for cur*2 < target {
				ans++
				cur *= 2
			}
		}

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1490/A
  Tags: greedy, math
  Rating: 800
*/
