// MIT License
//
// Copyright (c) 2018 Fris
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package GoStats

import (
	"log"
	"testing"
)

func init() {
	log.SetPrefix("[GoStats] ")
}

func TestSingular(t *testing.T) {
	if Singular(2) == true {
		log.Fatalln("Singular() didn't give the correct value!")
		t.Fail()
	}
}

func TestMax(t *testing.T) {
	if _, n := Max([]float64{1.1, 0.2}); n != 1.1 {
		log.Fatalln("Max() didn't give the correct value!")
		t.Fail()
	}
}

func TestMin(t *testing.T) {
	if _, n := Min([]float64{1.1, 0.2}); n != 0.2 {
		log.Fatalln("Min() didn't give the correct value!")
		t.Fail()
	}
}

func TestTotal(t *testing.T) {
	if Total([]float64{1.0, 0.5, 2.0}) != 3.5 {
		log.Fatalln("Total() didn't give the correct value!")
		t.Fail()
	}
}

func TestRange(t *testing.T) {
	if Range([]float64{0.1, 0.2}) != 0.1 {
		log.Fatalln("Range() didn't give the correct value!")
		t.Fail()
	}
}

func TestMean(t *testing.T) {
	if Mean([]float64{2.0, 2.0}) != 2 {
		log.Fatalln("Mean() didn't give the correct value!")
		t.Fail()
	}
}

func TestRepeats(t *testing.T) {
	if Repeats([]float64{0.1, 0.4, 0.4})[0.4] != 2 {
		log.Fatalln("Repeats() didn't give the correct value!")
		t.Fail()
	}
}

func TestMode(t *testing.T) {
	if mode, r := Mode([]float64{0.1, 0.4, 0.4}); r != 2 || mode != 0.4 {
		log.Fatalln("Mode() didn't give the correct value!")
		t.Fail()
	}
}

func TestMedian(t *testing.T) {
	if Median([]float64{0.5, 0.6, 0.4, 0.4}, true) != 0.45 {
		log.Fatalln("Median() didn't give the correct value! #1")
		t.Fail()
	}
	if Median([]float64{}, false) != 0 {
		log.Fatalln("Median(0) didn't give the correct value! #2")
		t.Fail()
	}
	if Median([]float64{0.5, 0.6, 0.4}, false) != 0.6 {
		log.Fatalln("Median() didn't give the correct value! #3")
		t.Fail()
	}
}

func TestVariance(t *testing.T) {
	if Variance([]float64{1.0, 2.0}) != 0.25 {
		log.Fatalln("Variance() didn't give the correct value!")
		t.Fail()
	}
}

func TestStandardDeviation(t *testing.T) {
	if StandardDeviation([]float64{10, 10}) != 0 {
		log.Fatalln("StandardDeviation() didn't give the correct value!")
		t.Fail()
	}
}