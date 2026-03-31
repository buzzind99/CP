//go:build codeforces_1896A

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
        possible := true
        n := nextInt()
        for i := range n {
            a := nextInt()
            if i == 0 && a != 1 { possible = false }
        }

        if possible {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1896/A
  Tags: sortings
  Rating: 800
*/
