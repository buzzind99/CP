//go:build codeforces_1551A

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
        n := nextInt()
        div := n/3
        if n%3 == 1 {
            fmt.Println(div+1, div)
        } else if n%3 == 2{
            fmt.Println(div, div+1)
        } else {
            fmt.Println(div, div)
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1551/A
  Tags: greedy, math
  Rating: 800
*/
