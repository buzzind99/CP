//go:build codeforces_1353B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "slices"
)

var sc = bufio.NewScanner(os.Stdin)

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

	t := nextInt()
	for range t {
        n, k := nextInt(), nextInt()
        a, b := make([]int, n), make([]int, n)
        sumA := 0
        for i := range n {
            a[i] = nextInt()
            sumA += a[i]
        }
        for i := range n {
            b[i] = nextInt()
        }
        slices.Sort(a)
        slices.Sort(b)
        for i := range k {
            if b[n-1-i] > a[i] {
                sumA += b[n-1-i] - a[i]
            }
        }

        fmt.Println(sumA)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1353/B
  Tags: greedy, sortings
  Rating: 800
*/
