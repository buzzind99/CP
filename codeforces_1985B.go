//go:build codeforces_1985B

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
        n := nextInt()
		if n == 3 {
			fmt.Println(3)
		} else {
			fmt.Println(2)
		}
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1985/B
  Tags: brute force, math, number theory
  Rating: 800
*/
