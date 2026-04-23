//go:build codeforces_1385B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
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

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		answer := []string{}
		for range 2*n {
			a := nextInt()

			if (!slices.Contains(answer, strconv.Itoa(a)) && len(answer) < n) {
				answer = append(answer, strconv.Itoa(a))
			}
		}

		fmt.Fprintln(wr, strings.Join(answer, " "))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1385/B
  Tags: greedy
  Rating: 800
*/
