//go:build codeforces_1619A

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
        if len(s)%2 != 0 {
            fmt.Println("NO")
        } else {
            half := len(s)/2
            if s[0:half] == s[half:] {
                fmt.Println("YES")
            } else {
                fmt.Println("NO")
            }
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1619/A
  Tags: implementation, strings
  Rating: 800
*/
