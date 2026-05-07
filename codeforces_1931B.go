//go:build codeforces_1931B

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
		arr := make([]int, 0, n)
		arrSum := 0.0
		for range n {
			a := nextInt()
			arr = append(arr, a)
			arrSum += float64(a)
		}
		arrAvg := arrSum/float64(n)

		currSum, count := 0.0, 0.0
		possible := true
		for i := n-1; i >= 0; i-- {
			currSum += float64(arr[i])
			count++
			currAvg := currSum/count
			if currAvg > arrAvg { possible = false }
		}

		if possible {
			fmt.Fprintln(wr, "YES")
		} else {
			fmt.Fprintln(wr, "NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1931/B
  Tags: greedy
  Rating: 800
*/
