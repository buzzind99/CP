//go:build codeforces_158A

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

	n, k := nextInt(), nextInt()
	scores := make([]int, n)
    for i := 0; i < n; i++ {
        scores[i] = nextInt()
    }
    
	sol := solve(n, k, scores)

	fmt.Println(sol)
}

func solve(n int, k int, scores []int) int {
    ans := 0
	cutoff := scores[k-1]
    for _, score := range scores {
        if score >= cutoff && score > 0 {
            ans++
        }
    }

    return ans
}

/*
  Link: https://codeforces.com/problemset/problem/158/A
  Tags: special problem, implementation
  Rating: 800
*/
