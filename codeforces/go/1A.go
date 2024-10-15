// Problem 1A:

// Theatre Square in the capital city of Berland has a rectangular shape with
// the size n*m meters. On the occasion of the city's anniversary, a decision
// was taken to pave the Square with square granite flagstones. Each flagstone
// is of the size a*a.

// What is the least number of flagstones needed to pave the Square? It's
// allowed to cover the surface larger than the Theatre Square, but the Square
// has to be covered. It's not allowed to break the flagstones. The sides of
// flagstones should be parallel to the sides of the Square.

// Input
// The input contains three positive integer numbers in the first line: n, m
// and a (0<=n, m, a<=109).

// Output
// Write the needed number of flagstones.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"math"
)

func CF1A(_r io.Reader, _w io.Writer){
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, a float64
	fmt.Fscan(in, &n, &m, &a)

	nof := math.Ceil(n/a) * math.Ceil(m/a)
	fmt.Fprint(out, int(nof))
}

func main() { CF1A(os.Stdin, os.Stdout) }
