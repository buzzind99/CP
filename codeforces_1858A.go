//go:build codeforces_1858A

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
        a, b, c := nextInt(), nextInt(), nextInt()
        if c%2 != 0 { a++ }

        if a > b {
            fmt.Println("First")
        } else {
            fmt.Println("Second")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1858/A
  Tags: games, greedy, math
  Rating: 800
*/
