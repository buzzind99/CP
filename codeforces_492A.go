//go:build codeforces_492A

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
    total, height := 0, 0
    for i := 1;; {
        cubes := i*(i+1)/2
        total += cubes
        if total <= n {
			height++
			i++
		} else {
			break
		}
    }

    fmt.Println(height)
}

/*
  Link: https://codeforces.com/problemset/problem/492/A
  Tags: implementation
  Rating: 800
*/
