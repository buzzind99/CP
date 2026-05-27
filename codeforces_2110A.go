//go:build codeforces_2110A

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
		arr := make([]int, 0, n)
		for range n {
			a := nextInt()
			arr = append(arr, a)
		}
		slices.Sort(arr)

		even := 0
		for i := range n {
			if isEven(arr[i]) { break }
			even++
		}
		for i := n-1; i >= 0; i-- {
			if isEven(arr[i]) { break }
			even++
		}
		odd := 0
		for i := range n {
			if !isEven(arr[i]) { break }
			odd++
		}
		for i := n-1; i >= 0; i-- {
			if !isEven(arr[i]) { break }
			odd++
		}

		fmt.Fprintln(wr, min(even, odd))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2110/A
  Tags: implementation, sortings
  Rating: 800
*/
