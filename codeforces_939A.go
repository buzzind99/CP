//go:build codeforces_939A

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

	n := nextInt()
	f := make([]int, 0, n+1)
	f = append(f, 0)
	for range n {
		a := nextInt()
		f = append(f, a)
	}

	for i := 1; i <= n; i++ {
		personA := i
		personB := f[personA]
		personC := f[personB]
		
		if f[personC] == personA {
			fmt.Println("YES")
			return
		}
	}

	fmt.Println("NO")
}

/*
  Link: https://codeforces.com/problemset/problem/939/A
  Tags: graphs
  Rating: 800
*/
