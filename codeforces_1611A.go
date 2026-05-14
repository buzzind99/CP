//go:build codeforces_1611A

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

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		s := next()
		firstDigit, lastDigit := int(s[0]-'0'), int(s[len(s)-1]-'0')
		evenCount := 0
		for i := range s {
			n := int(s[i]-'0')
			if isEven(n) { evenCount++ }
		}

		if isEven(lastDigit) {
			fmt.Fprintln(wr, 0)
		} else {
			if evenCount == 0 {
				fmt.Fprintln(wr, -1)
			} else {
				if isEven(firstDigit) {
					fmt.Fprintln(wr, 1)
				} else {
					fmt.Fprintln(wr, 2)
				}
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1611/A
  Tags: constructive algorithms, math
  Rating: 800
*/
