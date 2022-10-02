package main

// https://www.youtube.com/watch?v=9vhbKiwjJo8

import (
	"fmt"
	"math"
)

func main() {
	precision := 32
	whole := int64(math.Pow(2, float64(precision)))
	half := whole / 2
	quarter := whole / 4

	n := 1
	//X := []byte{'A', 'B'}
	r := []byte{2, 2}
	R := 0
	for i := 0; i <= n; i++ {
		R += int(r[i])
	}
	p := []float64{}
	for i := 0; i <= n; i++ {
		p = append(p, float64(r[i])/float64(R))
	}
	//EOF := 0
	x := "BABAB" + string([]byte{0})
	k := 5

	c := []byte{0}
	for j := 1; j <= n; j++ {
		sum := 0
		for i := 1; i <= j-1; j++ {
			sum += int(r[i])
		}
		c = append(c, byte(sum))
	}

	d := []byte{}
	for j := 0; j <= n; j++ {
		d = append(d, c[j]+r[j])
	}

	// algorithm starts
	s := 0
	// x is string input
	var a int64 = 0
	b := whole

	for i := 1; i <= k+1; i++ {
		w := b - a
		b = a + int64(math.Round(w*d[i]/R)) // change - d[i] wrong
		a = a + int64(math.Round(w*c[i]/R)) // change - c[i] wrong
		for b < half || a > half {
			if b < half {
				fmt.Println(0)
				for sidx := 0; sidx < s; sidx++ {
					fmt.Println(1)
				}
				s = 0
				a = 2 * a
				b = 2 * b
			} else if a > half {
				fmt.Println(1)
				for sidx := 0; sidx < s; sidx++ {
					fmt.Println(0)
				}
				s = 0
				a = 2 * (a - half)
				b = 2 * (b - half)
			}
		}
		for a > quarter && b < 3*quarter {
			s = s + 1
			a = 2 * (a - quarter)
			b = 2 * (b - quarter)
		}
	}

	s = s + 1
	if a <= quarter {
		fmt.Println(0)
		for i := 0; i < s; i++ {
			fmt.Println(1)
		}
	} else {
		fmt.Print(1)
		for i := 0; i < s; i++ {
			fmt.Println(0)
		}
	}

	fmt.Println("Hello Universe!")
}
