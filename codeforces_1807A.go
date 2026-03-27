//go:build codeforces_1807A

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
        if a+b == c {
            fmt.Println("+")
        } else {
            fmt.Println("-")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1807/A
  Tags: implementation
  Rating: 800
*/
