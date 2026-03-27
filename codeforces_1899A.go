//go:build codeforces_1899A

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
        if n%3 != 0 {
            fmt.Println("First")
        } else {
            fmt.Println("Second")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1899/A
  Tags: games, math, number theory
  Rating: 800
*/
