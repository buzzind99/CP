//go:build codeforces_978A

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
	arr := make([]int, 0, n)
	for range n {
		arr = append(arr, nextInt())
	}
	seen := make(map[int]struct{})
	result := []int{}
	
	for i := len(arr)-1; i >= 0; i-- {
		val := arr[i]
		if _, ok := seen[val]; !ok {
			seen[val] = struct{}{}
			result = append(result, val)
		}
	}

	fmt.Fprintln(wr, len(result))
	for i := len(result)-1; i >= 0; i-- {
		fmt.Fprint(wr, result[i], " ")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/978/A
  Tags: implementation
  Rating: 800
*/
