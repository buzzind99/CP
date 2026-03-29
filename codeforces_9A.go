//go:build codeforces_9A

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

func gcd(a int, b int) int {
    for b != 0 { a, b = b, a%b }
    return a
}

func main() {
	sc.Split(bufio.ScanWords)
    
	y, w := nextInt(), nextInt()
	highest := max(y, w)
	num := 7-highest
	den := 6

	if num == 6 {
        fmt.Println("1/1")
    } else if num == 0 {
        fmt.Println("0/1")
    } else {
        common := gcd(num, den)
        fmt.Printf("%d/%d\n", num/common, den/common)
    }
}

/*
  Link: https://codeforces.com/problemset/problem/9/A
  Tags: math, probabilities
  Rating: 800
*/
