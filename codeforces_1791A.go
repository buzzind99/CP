//go:build codeforces_1791A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "strings"
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
        c := next()
        if strings.Contains("cdefors", c) {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1791/A
  Tags: implementation, strings
  Rating: 800
*/
