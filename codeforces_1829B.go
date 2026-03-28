//go:build codeforces_1829B

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
        longest := -1
        counter := 0
        for range n {
            a := nextInt()
            if a == 0 { counter++ } else { counter = 0 }
            if counter > longest { longest = counter }
        }
        fmt.Println(longest)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1829/B
  Tags: implementation
  Rating: 800
*/
