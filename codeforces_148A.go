//go:build codeforces_148A

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

	k, l, m, n, d := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	if k == 1 {
		fmt.Println(d)
		return
	}

	damaged := make([]bool, d+1)
	count := 0
	for _, v := range []int{k, l, m, n} {
		for i := v; i <= d; i += v {
			if !damaged[i] {
				damaged[i] = true
				count++
			}
		}
	}

	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/148/A
  Tags: constructive algorithms, implementation, math
  Rating: 800
*/
