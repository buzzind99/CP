//go:build codeforces_1462A

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
			b := nextInt()
			arr = append(arr, b)
		}

		left, right := 0, n-1
		for i := range n {
			if i%2 == 0 {
				fmt.Fprint(wr, arr[left], " ")
				left++
			} else {
				fmt.Fprint(wr, arr[right], " ")
				right--
			}
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1462/A
  Tags: implementation, two pointers
  Rating: 800
*/
