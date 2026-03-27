//go:build codeforces_381A

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

func main() {
	sc.Split(bufio.ScanWords)
    
	n := nextInt()
	arr := make([]int, n)
	for i:= range n {
		arr[i] = nextInt()
	}

	sereja := 0
	dima := 0
	li := 0
	ri := n-1
	for i := range n {
		left := arr[li]
		right := arr[ri]
		if i%2 == 0 {
			if left > right {
				sereja += left
				li++
			} else {
				sereja += right
				ri--
			}
		} else {
			if left > right {
				dima += left
				li++
			} else {
				dima += right
				ri--
			}
		}
	}

	fmt.Printf("%d %d", sereja, dima)
}

/*
  Link: https://codeforces.com/problemset/problem/381/A
  Tags: greedy, implementation, two pointers
  Rating: 800
*/
