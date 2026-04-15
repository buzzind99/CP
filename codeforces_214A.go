//go:build codeforces_214A
 
package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m := nextInt(), nextInt()
	count := 0
	for a := 0; a*a <= n; a++ {
		b := n - a*a
		if a + b*b == m { count++ }
	}

	fmt.Fprintln(wr, count)
}
 
/*
  Link: https://codeforces.com/problemset/problem/214/A
  Tags: brute force
  Rating: 800
*/
