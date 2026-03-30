//go:build codeforces_1742B

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
        m := make(map[int]bool)
        possible := true
        for range n {
            a := nextInt()
            if !m[a] { m[a] = true } else { possible = false }
        }

        if possible {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1742/B
  Tags: greedy, implementation, sortings
  Rating: 800
*/
