//go:build codeforces_1834A

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
		plus, minus := 0, 0
		for range n {
			a := nextInt()
			if a == 1 { plus++ } else { minus++ }
		}

		ops := 0
		for minus%2 == 1 || minus > plus {
			ops++; minus--; plus++
		}

		fmt.Fprintln(wr, ops)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1834/A
  Tags: greedy, math
  Rating: 800
*/
