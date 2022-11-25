package main

import "math"

func circumference(r float64) float64 {
	return 2 * math.Pi * r
}

func vol_of_cuboid(l float64, b float64, h float64) float64 {
	return l * b * h
}