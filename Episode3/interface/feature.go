package main

import "math"

type Feature struct{}

type NewFeature interface {
	SqrtF(int) float64
	ExpF(int, int) int64
}

func NewF() *Feature {
	return &Feature{}
}

func (f *Feature) SqrtF(n int) float64 {
	return math.Sqrt(float64(n))
}

func (f *Feature) ExpF(x int, n int) int64 {
	return int64(math.Pow(float64(x), float64(n)))
}
