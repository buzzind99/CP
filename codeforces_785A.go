//go:build codeforces_785A

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
	poly := map[byte]int{'T': 4, 'C': 6, 'O': 8, 'D': 12, 'I': 20}
	sum := 0
	for range n {
		s := next()
		sum += poly[s[0]]
	}

	fmt.Println(sum)
}

/*
  Link: https://codeforces.com/problemset/problem/785/A
  Tags: implementation, strings
  Rating: 800
*/
