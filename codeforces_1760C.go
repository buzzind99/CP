//go:build codeforces_1760C

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
		
		sorted := slices.Clone(arr)
		slices.Sort(sorted)
		largestIdx, secondLargestIdx := len(arr)-1, len(arr)-2
		for _, v := range arr {
			if v == sorted[largestIdx] {
				fmt.Fprint(wr, v-sorted[secondLargestIdx], " ")
			} else {
				fmt.Fprint(wr, v-sorted[largestIdx], " ")
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
