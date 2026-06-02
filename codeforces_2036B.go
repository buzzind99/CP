//go:build codeforces_2036B

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
		n, k := nextInt(), nextInt()
		arr := make([]int, k+1)
		for range k {
			b, c := nextInt(), nextInt()
			arr[b] += c
		}
		slices.Sort(arr)

		total := 0
		amount := min(n, k)
		for i := 0; i < amount; i++ {
			total += arr[len(arr)-1-i]
		}

		fmt.Fprintln(wr, total)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2036/B
  Tags: greedy, sortings
  Rating: 800
*/
