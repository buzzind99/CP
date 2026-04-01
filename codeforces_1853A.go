//go:build codeforces_1853A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
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
		n := nextInt()
		sorted := true
		minVal := math.MaxInt
		var prev int
		for i := range n {
			a := nextInt()
			if i == 0 { prev = a; continue }
			if a < prev {
				sorted = false
			} else {
				if a-prev < minVal { minVal = a-prev }
			}
			prev = a
		}

		if !sorted {
			fmt.Println(0)
		} else {
			fmt.Println(minVal/2+1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1853/A
  Tags: brute force, greedy, math
  Rating: 800
*/
