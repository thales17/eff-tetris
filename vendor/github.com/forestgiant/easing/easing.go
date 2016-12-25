// Package easing is a go implementation of
// Robert Penner's Easing Equations - http://robertpenner.com/easing/
package easing

import (
	"math"
)

// BackIn easing
func BackIn(p float64) float64 {
	s := 1.70158
	return (p * p) * ((s+1)*p - s)
}

// BackOut easing
func BackOut(p float64) float64 {
	s := 1.70158
	p = p - 1
	return (p*p)*((s+1)*p+s) + 1
}

// BackInOut easing
func BackInOut(p float64) float64 {
	s := 1.70158
	s *= 1.525
	if p < .5 {
		return .5 * ((p * p) * ((s+1)*p - s))
	}

	p = p - 2
	return .5 * ((p*p)*((s+1)*p+s) + 2)
}

// BounceIn easing
func BounceIn(p float64) float64 {
	p = p - 1
	if p < (1 / 2.75) {
		return 1 - (7.5625 * p * p)
	} else if p < (2 / 2.75) {
		p -= 1.5 / 2.75
		return 1 - (7.5625*(p*p) + .75)
	} else if p < (2.5 / 2.75) {
		p -= 2.25 / 2.75
		return 1 - (7.5625*(p*p) + .9375)
	}

	p -= 2.625 / 2.75
	return 1 - (7.5625*(p*p) + .984375)
}

// BounceOut easing
func BounceOut(p float64) float64 {
	if p < (1 / 2.75) {
		return 7.5625 * p * p
	} else if p < (2 / 2.75) {
		p -= 1.5 / 2.75
		return 7.5625*(p*p) + .75
	} else if p < (2.5 / 2.75) {
		p -= 2.25 / 2.75
		return 7.5625*(p*p) + .9375
	}

	p -= 2.625 / 2.75
	return 7.5625*(p*p) + .984375
}

func BounceInOut(p float64) float64 {
	if p < 0.5 {
		p = p - 2
		if p < (1 / 2.75) {
			return .5 - (7.5625 * p * p)
		} else if p < (2 / 2.75) {
			p -= 1.5 / 2.75
			return .5 - (7.5625*(p*p) + .75)
		} else if p < (2.5 / 2.75) {
			p -= 2.25 / 2.75
			return .5 - (7.5625*(p*p) + .9375)
		}

		p -= 2.625 / 2.75
		return 1 - (7.5625*(p*p) + .984375)
	}

	p = p - 2
	if p < (1 / 2.75) {
		return 0.5 * (7.5625 * p * p)
	} else if p < (2 / 2.75) {
		p -= 1.5 / 2.75
		return 0.5 * (7.5625*(p*p) + .75)
	} else if p < (2.5 / 2.75) {
		p -= 2.25 / 2.75
		return 0.5 * (7.5625*(p*p) + .9375)
	}

	p -= 2.625 / 2.75
	return 0.5 * (7.5625*(p*p) + .984375)
}

// Linear easing
func Linear(p float64) float64 {
	return p
}

// QuadIn easing
func QuadIn(p float64) float64 {
	return p * p
}

// SineIn easing
func SineIn(p float64) float64 {
	return -1*math.Cos(p*(math.Pi/2)) + 1
}

// SineOut easing
func SineOut(p float64) float64 {
	return math.Sin(p * (math.Pi / 2))
}

// SineInOut easing
func SineInOut(p float64) float64 {
	// return 0.0
	return -0.5 * (math.Cos(math.Pi*p) - 1)
}
