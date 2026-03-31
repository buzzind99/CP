//go:build codeforces_1890A

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
        counts := make(map[int]int)
        for range n {
            counts[nextInt()]++
        }

        if len(counts) == 1 {
            fmt.Println("Yes")
        } else if len(counts) == 2 {
            vals := []int{}
            for _, c := range counts {
                vals = append(vals, c)
            }
            if abs(vals[0]-vals[1]) <= 1 {
                fmt.Println("Yes")
            } else {
                fmt.Println("No")
            }
        } else {
            fmt.Println("No")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1890/A
  Tags: constructive algorithms
  Rating: 800
*/
