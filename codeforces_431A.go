//go:build codeforces_431A

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
	const maxCapacity int = 100005 
    buf := make([]byte, maxCapacity)
    sc.Buffer(buf, maxCapacity)
	sc.Split(bufio.ScanWords)

	arr := [5]int{0, nextInt(), nextInt(), nextInt(), nextInt()}
	s := next()
	sum := 0
	for _, v := range s {
		n := int(v-'0') 
		sum += arr[n]
	}

    fmt.Println(sum)
}

/*
  Link: https://codeforces.com/problemset/problem/431/A
  Tags: implementation
  Rating: 800
*/
