//go:build codeforces_1760C

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
		largest, secondLargest := -1, -1
		for range n {
			x := nextInt()
			arr = append(arr, x)
			if x > largest {
				secondLargest = largest
				largest = x
			} else {
				secondLargest = max(secondLargest, x)
			}
		}

		for _, v := range arr {
			if v == largest {
				fmt.Fprint(wr, v-secondLargest, " ")
			} else {
				fmt.Fprint(wr, v-largest, " ")
			}
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1760/C
  Tags: data structures, implementation, sortings
  Rating: 800
*/
