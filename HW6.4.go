package HW6.4

import "math"


func EquationSquare(a, b, c float64) (x1, x2, d float64) {
	
	if a == 0 {
		
		x1 = math.Sqrt(-1)
		x2, d = x1, x1
		return x1, x2, d
	}
	
	d = b*b - 4*a*c

	if d == 0 { /
		x1 = -b / (2 * a) 
		x2 = x1
	} else if d > 0 { 
		a2 := 2 * a
		ds := math.Sqrt(d) 
		x1 = (-b - ds) / a2
		x2 = (-b + ds) / a2
	}
	
	return x1, x2, d
}