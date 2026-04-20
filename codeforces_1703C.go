//go:build codeforces_1703C

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
		digits := make([]int, 0, n)
		for range n {
			digits = append(digits, nextInt())
		}
		for i := range n {
			nextInt()
			sequence := next()
			ops := 0
			for _, char := range sequence {
				if char == 'D' { ops++ } else { ops-- }
			}
			digits[i] = (10+digits[i]+ops)%10
		}

		for _, v := range digits {
			fmt.Fprint(wr, v, " ")
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1703/C
  Tags: brute force, implementation, strings
  Rating: 800
*/
