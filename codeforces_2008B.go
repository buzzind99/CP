//go:build codeforces_2008B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
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
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, s := nextInt(), next()
		r := int(math.Round(math.Sqrt(float64(n))))
		if r * r != n {
			fmt.Println("NO")
			continue
		}

		isValid := true
		for i := 0; i < r; i++ {
			for j := 0; j < r; j++ {
				idx := i*r+j
				isBoundary := (i == 0 || i == r-1 || j == 0 || j == r-1)
				if isBoundary && s[idx] != '1' {
					isValid = false
					break
				}
				if !isBoundary && s[idx] != '0' {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}

		if isValid {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2008/B
  Tags: brute force, math, strings
  Rating: 800
*/
