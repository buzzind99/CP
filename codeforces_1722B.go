//go:build codeforces_1722B

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
        x, y := next(), next()
        same := true
        for i := range n {
            if x[i] != y[i] {
                if x[i] == 'R' || y[i] == 'R' {
                    same = false
                    break
                }
            }
        }

        if same {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1722/B
  Tags: implementation
  Rating: 800
*/
