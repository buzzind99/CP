//go:build codeforces_1862A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, m := nextInt(), nextInt()
		carpet := make([]string, 0, n)
		for range n {
			carpet = append(carpet, next())
		}

		target := "vika"
		count := 0
		for j := range m {
			for i := range n {
				if count < 4 && carpet[i][j] == target[count] { count++; break }
			}
		}

		if count == 4 {
			fmt.Fprintln(wr, "YES")
		} else {
			fmt.Fprintln(wr, "NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1862/A
  Tags: dp, greedy, implementation, strings
  Rating: 800
*/
