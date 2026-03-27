//go:build codeforces_1703A

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
        s:= strings.ToLower(next())
        if s == "yes" {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1703/A
  Tags: implementation
  Rating: 800
*/
