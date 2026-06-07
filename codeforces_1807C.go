//go:build codeforces_1807C

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
		n, s := nextInt(), next()
		posParity := make([]int, 26)
		for i := range 26 {
			posParity[i] = -1
		}

		possible := true
		for i := range n {
			charIdx := int(s[i] - 'a')
			currentParity := i%2
			if posParity[charIdx] == -1 {
				posParity[charIdx] = currentParity
			} else if posParity[charIdx] != currentParity {
				possible = false
				break
			}
		}

		if possible {
			fmt.Fprintln(wr, "YES")
		} else {
			fmt.Fprintln(wr, "NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1807/C
  Tags: greedy, implementation, strings
  Rating: 800
*/
