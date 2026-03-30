//go:build codeforces_1971A

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
		x, y := nextInt(), nextInt()
		if x <= y {
			fmt.Println(x, y)
		} else {
			fmt.Println(y, x)
		}
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1971/A
  Tags: implementation, sortings
  Rating: 800
*/
