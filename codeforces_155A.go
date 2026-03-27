//go:build codeforces_155A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	n := nextInt()
	max := 0
	min := 0
	count := 0
	for i := range n {
		x := nextInt()
		if i == 0 {
			max, min = x, x
			continue
		}
		if x > max {
			max = x
			count++
		} else if x < min {
			min = x
			count++
		}
	}

	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/155/A
  Tags: brute force
  Rating: 800
*/
