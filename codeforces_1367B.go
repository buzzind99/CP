//go:build codeforces_1367B

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

func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func main() {
	sc.Split(bufio.ScanWords)

	t := nextInt()
	for range t {
        n := nextInt()
		odd, even := 0, 0
		for i := range n {
			a := nextInt()
			if a%2 == 0 && i%2 != 0 { odd++ }
			if a%2 == 1 && i%2 != 1 { even++ }
		}

		if odd == even {
			fmt.Println(odd)
		} else {
			fmt.Println(-1)
		}
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1367/B
  Tags: greedy, math
  Rating: 800
*/
