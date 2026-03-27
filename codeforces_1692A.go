//go:build codeforces_1692A

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
    count := 0
    for range t {
        a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()
        if b > a { count++ }
        if c > a { count++ }
        if d > a { count++ }

        fmt.Println(count)
        count = 0
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1692/A
  Tags: implementation
  Rating: 800
*/
