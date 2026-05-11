//go:build codeforces_2167C

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

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		odd, even := 0, 0
		arr := make([]int, 0, n)
		for range n {
			a := nextInt()
			arr = append(arr, a)
			if isEven(a) { even++ } else { odd++ }
		}

		if even != 0 && odd != 0 {
			slices.Sort(arr)
		}

		for _, v := range arr {
			fmt.Fprint(wr, v, " ")
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2167/C
  Tags: constructive algorithms, greedy, implementation, sortings
  Rating: 800
*/
