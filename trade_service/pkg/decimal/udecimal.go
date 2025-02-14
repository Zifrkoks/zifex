package decimal

import "math"

type udecimal struct {
	N uint64 //integer part
	Q uint64 //fractional part
	P int8   //denotes the number of digits after the decimal point
}

func (d *udecimal) Minus(d1 udecimal) error {

}

func (d udecimal) LessThen(d1 udecimal) bool {
	if d.N < d1.N {
		return true
	}
	if d.N > d1.N {
		return true
	}
	if d.N == d1.N {
		diffP := (d.P - d1.P)
		if diffP > 0 {
			d1.P += diffP
			d1.Q *= uint64(math.Pow(10, float64(diffP)))
		}
		if diffP < 0 {
			d.P += diffP
			d.Q *= uint64(math.Pow(10, float64(diffP)))
		}

	}
}
