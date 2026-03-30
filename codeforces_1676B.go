//go:build codeforces_1676B

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
        minVal := 10000001
        for i := range n {
            a := nextInt()
            arr[i] = a
            if a < minVal { minVal = a }
        }

        total := 0
        for _, v := range arr {
            total += v-minVal
        }

        fmt.Println(total)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1676/B
  Tags: brute force, math
  Rating: 800
*/
