//go:build codeforces_2091B

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
		n, x := nextInt(), nextInt()
		arr := make([]int, 0, n)
		for range n {
			a := nextInt()
			arr = append(arr, a)
		}
		slices.Sort(arr)
		slices.Reverse(arr)

		strongTeams := 0
		currentTeamSize := 0
		for i := 0; i < n; i++ {
			currentTeamSize++
			neededSize := (x+arr[i]-1)/arr[i]
			if currentTeamSize >= neededSize {
				strongTeams++
				currentTeamSize = 0
			}
		}

		fmt.Fprintln(wr, strongTeams)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2091/B
  Tags: dp, greedy, sortings
  Rating: 800
*/
