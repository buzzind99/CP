//go:build codeforces_1873B

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

	t := nextInt()
    for range t {
        n := nextInt()
        arr := make([]int, n)
        minVal, minIdx := 10, -1
        for i := range n {
            a := nextInt()
            arr[i] = a
            if a < minVal { minVal, minIdx = a, i }
        }
        arr[minIdx]++
        total := 1
        for _, v := range arr {
            total *= v
        }

        fmt.Println(total)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1873/B
  Tags: brute force, greedy, math
  Rating: 800
*/
