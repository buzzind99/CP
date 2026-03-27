//go:build codeforces_1669A

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
        switch n := nextInt(); {
        case n >= 1900:
            fmt.Println("Division 1")
        case n >= 1600:
            fmt.Println("Division 2")
        case n >= 1400:
            fmt.Println("Division 3")
        default:
            fmt.Println("Division 4")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1669/A
  Tags: implementation
  Rating: 800
*/
