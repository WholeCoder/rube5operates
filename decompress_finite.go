package main

// https://www.youtube.com/watch?v=RFPwltOz1IU&t=4s

import (
	"fmt"
	"math"
)

func main() {
	precision := 32
	whole := int64(math.Pow(2, float64(precision)))
	half := whole / 2
	quarter := whole / 4

	a := 0
	b := whole
	z := 0
	i := 1

	x := "0001"
	M := len(x)
	for i <= precision && i <= M {
		if x[i] == '0' {
			z = z + int64(math.Pow(2, float64(precision-i)))
		}
		i++
	}
	for true {
		for j := 0; j <= n; j++ {
			w := b - a
			b0 = a + int64(math.Round(w*d[j]/R))
			a0 = a + int64(math.Round(w*c[j]/R))
			if z >= a0 && z < b0 {
				fmt.Println(j)
				a = a0
				b = b0
				if j == 0 {
					return // quit? -> from algorithm
				}
			}
		}
		for b < half || a > half {
			if b < half {
				a = 2 * a
				b = 2 * b
				z = 2 * z
			} else if a > half {
				a = 2 * (a - half)
				b = 2 * (b - half)
				z = 2 * (z - half)
			}
			if i <= M && x[i] == '1' {
				z = z + 1
			}
			i++
		}
		for a > quarter && b < 3*quarter {
			a = 2 * (a - quarter)
			b = 2 * (b - quarter)
			z = 2 * (z - quarter)
			if i <= M && x[i] == '1' {
				z++
			}
			i++
		}
	}
}
