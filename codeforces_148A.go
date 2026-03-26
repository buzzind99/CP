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

	arr := []int{k, l, m, n}
	set := make(map[int]struct{})
	for _, v := range arr {
		i := 1
		for {
			tmp := i*v
			if tmp > d { break }
			set[tmp] = struct{}{}
			i++
		}
	}

	fmt.Println(len(set))
}

/*
  Link: https://codeforces.com/problemset/problem/148/A
  Tags: constructive algorithms, implementation, math
  Rating: 800
*/
