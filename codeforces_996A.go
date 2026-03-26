//go:build codeforces_996A

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
	count := 0
	for n > 0 {
		if n >= 100 {
			count += n/100
			n %= 100
		}
		if n >= 20 {
			count += n/20
			n %= 20
		}
		if n >= 10 {
			count += n/10
			n %= 10
		}
		if n >= 5 {
			count += n/5
			n %= 5
		}
		if n >= 1 {
			count += n
			break
		}
	}

	fmt.Println(count)
}

/*
  Link: https://codeforces.com/problemset/problem/996/A
  Tags: implementation
  Rating: 800
*/
