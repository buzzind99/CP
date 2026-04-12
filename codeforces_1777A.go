//go:build codeforces_1777A

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
		for range n {
			arr = append(arr, nextInt())
		}

		count := 0
		for i := 0; i < n-1; i++ {
			if arr[i]%2 == arr[i+1]%2 { count++ }
		}

		fmt.Fprintln(wr, count)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1777/A
  Tags: greedy, math
  Rating: 800
*/
