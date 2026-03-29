//go:build codeforces_1829A

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
    target := "codeforces"
    for range t {
        s := next()
        count := 0
        for i := range 10 {
            if s[i] != target[i] { count++ }
        }
        fmt.Println(count)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1829/A
  Tags: implementation, strings
  Rating: 800
*/
