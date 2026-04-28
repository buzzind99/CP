//go:build codeforces_1974A

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
		x, y := nextInt(), nextInt()
		ans := (y+1)/2

		totalCellsInScreens := ans*15
		usedByY := y*4
		remainingSlots := totalCellsInScreens-usedByY
		if x > remainingSlots {
			extraX := x-remainingSlots
			ans += (extraX+14)/15
		}

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1974/A
  Tags: greedy, math
  Rating: 800
*/
