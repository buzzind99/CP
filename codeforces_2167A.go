//go:build codeforces_2167A

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

	t := nextInt()
    for range t {
        possible := true
		prev := nextInt()
		for range 3 {
			curr := nextInt()
			if prev != curr { possible = false }
			prev = curr
		}

		if possible {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
    }
}

/*
  Link: https://codeforces.com/problemset/problem/2167/A
  Tags: math, sortings
  Rating: 800
*/
