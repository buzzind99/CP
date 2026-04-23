//go:build codeforces_2051B

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
		n, a, b, c := nextInt(), nextInt(), nextInt(), nextInt()
		arr := [3]int{a,b,c}
		total := a+b+c
		nDiv := n/total
		sum := nDiv*total
		remainder, count := n-sum, nDiv*3
		if remainder > 0 {
			for i := 0; sum < n; i++ {
				sum += arr[i%3]
				count++
			}
		}

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2051/B
  Tags: binary search, math
  Rating: 800
*/
