//go:build codeforces_1999A

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
        a, b := int(s[0]-'0'), int(s[1]-'0') 
        fmt.Println(a+b)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1999/A
  Tags: implementation, math
  Rating: 800
*/
