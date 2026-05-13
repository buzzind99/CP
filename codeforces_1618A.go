//go:build codeforces_1618A

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
		arr := [7]int{}
		for i := range 7 {
			arr[i] = nextInt()
		}
		firstThreeSum := arr[0]+arr[1]+arr[2]

		if arr[6] == firstThreeSum {
			fmt.Fprintln(wr, arr[0], arr[1], arr[2])
		} else {
			fmt.Fprintln(wr, arr[0], arr[1], arr[3])
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1618/A
  Tags: math, sortings
  Rating: 800
*/
