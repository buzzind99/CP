//go:build codeforces_1772A

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
		s := next()
        fmt.Println(int(s[0]-'0')+int(s[2]-'0'))
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1772/A
  Tags: implementation
  Rating: 800
*/
