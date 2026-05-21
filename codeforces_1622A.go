//go:build codeforces_1622A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
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
		sticks := make([]int, 0, 3)
		for range 3 { sticks = append(sticks, nextInt()) }
		slices.Sort(sticks)
		l1, l2, l3 := sticks[0], sticks[1], sticks[2]

		if l1+l2 == l3 || (l1 == l2 && l3%2 == 0) || (l2 == l3 && l1%2 == 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1622/A
  Tags: geometry, math
  Rating: 800
*/
