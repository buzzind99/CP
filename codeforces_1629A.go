//go:build codeforces_1629A

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

type Software struct {
	req int
	gain int
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, k := nextInt(), nextInt()
		reqs := make([]int, n)
		for i := range n {
			reqs[i] = nextInt()
		}
		softwares := make([]Software, n)
		for i := range n {
			softwares[i] = Software{
				req:  reqs[i],
				gain: nextInt(),
			}
		}
		slices.SortFunc(softwares, func(a, b Software) int {
			return a.req - b.req
		})

		currentRAM := k
		for _, soft := range softwares {
			if currentRAM >= soft.req {
				currentRAM += soft.gain
			} else {
				break
			}
		}

		fmt.Fprintln(wr, currentRAM)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1629/A
  Tags: brute force, greedy, sortings
  Rating: 800
*/
