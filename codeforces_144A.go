//go:build codeforces_144A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
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

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	arr := make([]int, n)
	for i := range n {
		arr[i] = nextInt()
	}

	min := slices.Min(arr)
	max := slices.Max(arr)
	count := 0
	for {
		if arr[0] == max && arr[n-1] == min { break }
		for i := 0; i < n; i++ {
			if arr[i] == max && i != 0 {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				count++
				break
			}
			if arr[i] == min && i != n-1 {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				count++
				break
			}
		}
	}

	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/144/A
  Tags: implementation
  Rating: 800
*/
