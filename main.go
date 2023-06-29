package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	data := ms120 // ms120, ms200, ms400
	Hₒ := 3.0

	xBar := stat.Mean(data, nil)
	σ := stat.StdDev(data, nil)

	n := float64(len(data))
	t := (xBar - Hₒ) / (σ / math.Sqrt(n))

	// Degrees of freedom
	df := n - 1

	dist := distuv.StudentsT{Mu: 0, Sigma: 1, Nu: df}

	p := 1 - dist.CDF(t)

	if p < 0.05 {
		fmt.Printf("Reject Hₒ, we have significant evidence to suggest that reach exceeds 3.0 (p=%.10f)", p)
	} else {
		fmt.Printf("Failed to reject Hₒ, no significant evidence that reach exceeds 3.0 (p=%.10f)", p)
	}
}
