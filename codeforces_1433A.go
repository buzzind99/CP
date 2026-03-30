//go:build codeforces_1433A

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
    arr := [5]int{0,1,3,6,10}
	for range t {
        x := next()
        firstDigit := int(x[0]-'0')
        numOfDigit := len(x)
        fmt.Println((firstDigit-1)*10+arr[numOfDigit])
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1433/A
  Tags: implementation, math
  Rating: 800
*/
