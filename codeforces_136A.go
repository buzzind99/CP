//go:build codeforces_136A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	n := nextInt()
	arr := make([]int, n)
	for i := range n {
		p := nextInt()
		arr[p-1] = i+1
	}
	strArr := make([]string, n)
	for i, v := range arr {
		strArr[i] = strconv.Itoa(v)
	}
	fmt.Println(strings.Join(strArr, " "))
}

/*
  Link: https://codeforces.com/problemset/problem/136/A
  Tags: implementation
  Rating: 800
*/
