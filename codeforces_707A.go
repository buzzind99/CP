//go:build codeforces_707A

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
 
	n, m := nextInt(), nextInt()
	colored := false
	for range n {
		for range m {
			color := next()
			if color[0] == 'C' || color[0] == 'M' || color[0] == 'Y' { colored = true }
		}
	}

	if colored {
		fmt.Println("#Color")
	} else {
		fmt.Println("#Black&White")
	}
}

/*
  Link: https://codeforces.com/problemset/problem/707/A
  Tags: implementation
  Rating: 800
*/