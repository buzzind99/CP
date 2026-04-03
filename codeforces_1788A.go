//go:build codeforces_1788A

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
		arr := [3]int{0,0,0}
		nums := make([]int, 0, n)
		for range n {
			a := nextInt()
			arr[a]++
			nums = append(nums, a)
		}

		if arr[2] == 0 {
			fmt.Fprintln(wr, 1)
		} else if arr[2]%2 != 0 {
			fmt.Fprintln(wr, -1)
		} else {
			twoCount := 0
			for i, v := range nums {
				if v == 2 { twoCount++ }
				if twoCount == arr[2]/2 {
					fmt.Fprintln(wr, i+1)
					break
				}
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1788/A
  Tags: brute force, implementation, math
  Rating: 800
*/
