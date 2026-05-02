//go:build codeforces_1800A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func abs(x int) int {
	if x < 0 { return -x }
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, s := nextInt(), strings.ToLower(next())
		compressed := make([]byte, 0, 4)
		for i := 0; i < n; i++ {
			if i == 0 || s[i] != s[i-1] {
				compressed = append(compressed, s[i])
			}
		}

		if string(compressed) == "meow" {
			fmt.Fprintln(wr, "YES")
		} else {
			fmt.Fprintln(wr, "NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1800/A
  Tags: implementation, strings
  Rating: 800
*/
