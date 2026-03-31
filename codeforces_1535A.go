//go:build codeforces_1535A

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
        a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()

        if max(a, b) < min(c, d) || max(c, d) < min(a, b) {
            fmt.Println("NO")
        } else {
            fmt.Println("YES")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1535/A
  Tags: brute force, implementation
  Rating: 800
*/
