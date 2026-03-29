//go:build codeforces_1343B

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
        if (n/2)%2 == 0 {
            sumEven := 0
            sumOdd := 0
            fmt.Println("YES")
            for i := range n/2 {
                n := (i+1)*2
                sumEven += n
                fmt.Print(n, " ")
            }
            for i := range n/2-1 {
                n := 1+i*2
                sumOdd += n
                fmt.Print(n, " ")
            }
            fmt.Println(sumEven-sumOdd)
        } else {
            fmt.Println("NO")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1343/B
  Tags: constructive algorithms, math
  Rating: 800
*/
