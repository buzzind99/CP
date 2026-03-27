//go:build codeforces_1676A

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
        sumLeft := 0
        sumRight := 0
        for i := 0; i+3 < 6; i++ {
            sumLeft += int(s[i])
            sumRight += int(s[i+3])
        }
        if sumLeft == sumRight {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1676/A
  Tags: implementation
  Rating: 800
*/
