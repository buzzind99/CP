//go:build codeforces_1537A

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
		sum := 0
		count := 0
		for range n {
			sum += nextInt()
			count++
		}

		if sum > count {
			fmt.Fprintln(wr, sum-count)
		} else if sum == count {
			fmt.Fprintln(wr, 0)
		} else {
			fmt.Fprintln(wr, 1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1537/A
  Tags: greedy, math
  Rating: 800
*/
