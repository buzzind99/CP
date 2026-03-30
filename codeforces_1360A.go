//go:build codeforces_1360A

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
        if a > b { a, b = b, a }
        side := a*2
        if b > side { side = b }

        fmt.Println(side*side)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1360/A
  Tags: greedy, math
  Rating: 800
  Notes: rectangle area inside of a rectangle
*/
