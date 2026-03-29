//go:build codeforces_490A

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

	n := nextInt()
	var arr [4][]int 
    for i := 1; i <= n; i++ {
        skill := nextInt()
        arr[skill] = append(arr[skill], i)
    }

    numTeams := min(len(arr[1]), len(arr[2]), len(arr[3]))
    fmt.Println(numTeams)
    for i := range numTeams {
        fmt.Printf("%d %d %d\n", arr[1][i], arr[2][i], arr[3][i])
    }
}

/*
  Link: https://codeforces.com/problemset/problem/490/A
  Tags: greedy, implementation, sortings
  Rating: 800
*/
