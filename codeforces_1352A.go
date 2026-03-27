//go:build codeforces_1352A

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

	t := nextInt()
	for range t {
        n := nextInt()
        var sol []int
        multiplier := 1
        for n > 0 {
            remainder := n%10
            if remainder > 0 {
                sol = append(sol, remainder * multiplier)
            }
            n /= 10
            multiplier *= 10
        }

        fmt.Println(len(sol))
        for i, v := range sol {
            fmt.Printf("%d ", v)
            if i == len(sol)-1 { fmt.Println() }
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1352/A
  Tags: implementation, math
  Rating: 800
*/
