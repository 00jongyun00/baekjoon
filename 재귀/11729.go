package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func hanoi(n int, from int, mid int, to int, wr *bufio.Writer) {
	if n == 1 {
		//fmt.Fprintf(wr, "원반 %d을 %d에서 %d로 이동\n", n, from, to)
		fmt.Fprintf(wr, "%d %d\n", from, to)
		return
	}
	// 시작 -> 경유지 -> 도착
	hanoi(n-1, from, to, mid, wr)
	//fmt.Fprintf(wr, "원반 %d를 %d에서 %d로 이동\n", n, from, to)
	fmt.Fprintf(wr, "%d %d\n", from, to)
	hanoi(n-1, mid, from, to, wr)
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanf(rd, "%d\n", &n)
	fmt.Fprintf(wr, "%d\n", int(math.Pow(2, float64(n)))-1)
	hanoi(n, 1, 2, 3, wr)

	wr.Flush()
}
