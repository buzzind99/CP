//go:build codeforces_1472B

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

func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func main() {
	sc.Split(bufio.ScanWords)

	t := nextInt()
	for range t {
        n := nextInt()
		one, two := 0, 0
		for range n {
			a := nextInt()
			if a == 1 { one++ } else { two++ }
		}

		total := one + 2*two
		if total%2 != 0 {
			fmt.Println("NO")
			continue
		}

		half := total / 2
        neededFromOnes := half
        if (half / 2) <= two {
            neededFromOnes = half%2
        } else {
            neededFromOnes = half-(two*2)
        }

        if one >= neededFromOnes {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1472/B
  Tags: dp, greedy, math
  Rating: 800
*/
