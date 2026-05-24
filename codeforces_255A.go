//go:build codeforces_255A

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

	n := nextInt()
	chest, biceps, back := 0, 0, 0
	for i := range n {
		a := nextInt()
		switch i%3 {
		case 0: chest += a
		case 1: biceps += a
		default: back += a
		}
	}

	if chest > biceps && chest > back {
		fmt.Fprintln(wr, "chest")
	} else if biceps > chest && biceps > back {
		fmt.Fprintln(wr, "biceps")
	} else {
		fmt.Fprintln(wr, "back")
	}

}

/*
  Link: https://codeforces.com/problemset/problem/255/A
  Tags: implementation
  Rating: 800
*/
