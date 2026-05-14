//go:build codeforces_1789A

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

func gcd(a int, b int) int {
	for b != 0 { a, b = b, a%b }
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		arr := make([]int, 0, n)
		for range n {
			arr = append(arr, nextInt())
		}

		solved := false
		for i := 0; i < n; i++ {
			if solved { break }
			for j := i + 1; j < n; j++ {
				if gcd(arr[i], arr[j]) <= 2 {
					fmt.Fprintln(wr, "Yes")
					solved = true
					break
				}
			}
		}

		if !solved {
			fmt.Fprintln(wr, "No")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1789/A
  Tags: brute force, math, number theory
  Rating: 800
*/
