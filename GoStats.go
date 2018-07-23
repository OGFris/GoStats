/**
* MIT License
*
* Copyright (c) 2018 Fris
*
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
*/

package GoStats

import (
	"math"
	"sort"
)

const (
	MaxFloat64 = math.MaxFloat64
	MinFloat64 = -(MaxFloat64)
)

//Get the Mean of the float64 numbers.
func Mean(numbers []float64) float64 {
	return Total(numbers) / float64(len(numbers))
}

//Get the Mode of the float64 numbers.
func Mode(numbers []float64) (mode float64, r int) {
	r = math.MinInt64
	for i,n := range Repeats(numbers) {
		if n > r {
			mode = i
			r = n
		}
	}
	return
}

//Get the Repeats of the float64 numbers.
func Repeats(numbers []float64) (r map[float64]int) {
	r = make(map[float64]int)
	for _,n := range numbers {
		r[n]++
	}
	return
}

//Get the Total of the float64 numbers.
func Total(numbers []float64) (total float64) {
	for _,n := range numbers {
		total += n
	}
	return
}

//Get the Median of the float64 numbers.
func Median(numbers []float64, sort bool) float64 {
	if sort {
		numbers = Sort(numbers)
	}
	if len(numbers) == 0 {
		return 0
	}
	if Singular(float64(len(numbers))) {
		return numbers[len(numbers) / 2]
	} else {
		return (numbers[len(numbers) / 2] + numbers[len(numbers) / 2 - 1]) / 2
	}
}

//Get whether a number is singular or not.
func Singular(number float64) bool {
	n := number / 2
	return float64(int(n)) != n
}

//Get the smallest number of the float64 numbers.
func Min(numbers []float64) (iMin int,min float64) {
	min = MaxFloat64
	for i,n := range numbers {
		if n < min {
			min = n
			iMin = i
		}
	}
	return
}

//Get the biggest number of the float64 numbers.
func Max(numbers []float64) (iMax int, max float64) {
	max = MinFloat64
	for i,n := range numbers {
		if n > max {
			max = n
			iMax = i
		}
	}
	return
}

//Get the Range of the float64 numbers.
func Range(numbers []float64) float64 {
	_,max := Max(numbers)
	_,min := Min(numbers)
	return max - min
}

//Get the Variance of the float64 numbers.
func Variance(numbers []float64) (variance float64) {
	mean := Mean(numbers)
	for _,n := range numbers {
		dif := Range([]float64{mean,n})
		variance += dif * dif
	}
	variance /= float64(len(numbers))
	return
}

//Get the Standard Deviation of the float64 numbers.
func StandardDeviation(numbers []float64) float64 {
	return math.Sqrt(Variance(numbers))
}

//Get the sorted []float64.
func Sort(numbers []float64) []float64 {
	sort.Float64s(numbers)
	return numbers
}
