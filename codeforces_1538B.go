//go:build codeforces_1538B

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
		arr := make([]int, 0, n)
		sum := 0
		for range n {
			a := nextInt()
			arr = append(arr, a)
			sum += a
		}

		if sum%n != 0 {
			fmt.Fprintln(wr, -1)
			continue
		}

		avg := sum/n
		count := 0
		for _, v := range arr {
			if v > avg { count++ }
		}

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1538/B
  Tags: greedy, math
  Rating: 800
*/
