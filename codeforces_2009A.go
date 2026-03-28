//go:build codeforces_2009A

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
        a, b := nextInt(), nextInt()
		fmt.Println(b-a)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/2009/A
  Tags: brute force, math
  Rating: 800
*/
