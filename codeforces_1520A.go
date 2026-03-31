//go:build codeforces_1520A

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
        n, s := nextInt(), next()
        arr := make([]bool, 26)
        curr := int(s[0]-'A')
        arr[curr] = true
        distracted := false
        for i := 1; i < n; i++ {
            task := int(s[i]-'A')
            if !arr[task] {
                arr[task] = true
            } else {
                if task != curr {
                    distracted = true
                    break
                }
            }
            curr = task
        }

        if !distracted {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}

/*
  Link: https://codeforces.com/problemset/problem/1520/A
  Tags: brute force, implementation
  Rating: 800
*/
