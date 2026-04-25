//go:build codeforces_2091A

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
		n := nextInt()
		m := make(map[int]int)
		target := map[int]int{0:3,1:1,2:2,3:1,4:0,5:1}
		matched, ans := 0, 0
		for i := range n {
			a := nextInt()
			if m[a] < target[a] {
				m[a]++
				matched++
			}

			if matched == 8 && ans == 0 {
				ans = i+1
			}
		}

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2091/A
  Tags: greedy, strings
  Rating: 800
*/
