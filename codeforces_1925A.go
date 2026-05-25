//go:build codeforces_1925A

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
		n, k := nextInt(), nextInt()
		chars := make([]byte, 0, k)
		for i := range k {
			chars = append(chars, byte(int('a')+i))
		}
		str := string(chars)
		slices.Reverse(chars)
		reversed := string(chars)
		ans := str
		for i := range n-1 {
			if i%2 == 0 { ans += reversed } else { ans += str }
		}

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1925/A
  Tags: constructive algorithms, greedy, strings
  Rating: 800
*/
