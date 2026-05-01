//go:build codeforces_1675A

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

func isEven (n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		a, b, c, dogs, cats := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
		availableDogFood := min(dogs, a)
		availableCatFood := min(cats, b)
		missingDogFood := dogs-availableDogFood
		missingCatFood := cats-availableCatFood
		
		if missingDogFood+missingCatFood > c {
			fmt.Fprintln(wr, "NO")
		} else {
			fmt.Fprintln(wr, "YES")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1675/A
  Tags: greedy, math
  Rating: 800
*/
