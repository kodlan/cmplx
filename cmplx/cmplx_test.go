package cmplx

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func almostEqual[T float32 | float64](a, b T) bool {
	if diff := cmp.Diff(a, b, cmpopts.EquateApprox(1e-9, 0)); diff != "" {
		return true
	}
	return false
}

func TestComplexAdd(t *testing.T) {
	var c1 = Complex{1, 1}
	var c2 = Complex{2, 2}
	var c3 = Add(&c1, &c2)

	var expected = complex(1, 1) + complex(2, 2)

	if almostEqual(float64(c3.im), imag(expected)) && almostEqual(float64(c3.re), real(expected)) {
		t.Errorf("%v + %v should be %v, got %v instead", c1, c2, expected, c3)
	}
}

func TestComplexAddMethod(t *testing.T) {
	var c1 = Complex{1, 1}
	var c2 = Complex{2, 2}
	c1.Add(&c2)

	var expected = complex(1, 1) + complex(2, 2)

	if almostEqual(float64(c1.im), imag(expected)) && almostEqual(float64(c1.re), real(expected)) {
		t.Errorf("%v + %v should be %v, got %v instead", c1, c2, expected, c1)
	}
}

func TestComplexSub(t *testing.T) {
	var c1 = Complex{1, 1}
	var c2 = Complex{2, 2}
	var c3 = Sub(&c1, &c2)

	var expected = complex(1, 1) - complex(2, 2)

	if almostEqual(float64(c3.im), imag(expected)) && almostEqual(float64(c3.re), real(expected)) {
		t.Errorf("%v - %v should be %v, got %v instead", c1, c2, expected, c3)
	}
}

func TestComplexSubMethod(t *testing.T) {
	var c1 = Complex{1, 1}
	var c2 = Complex{2, 2}
	c1.Sub(&c2)

	var expected = complex(1, 1) - complex(2, 2)

	if almostEqual(float64(c1.im), imag(expected)) && almostEqual(float64(c1.re), real(expected)) {
		t.Errorf("%v - %v should be %v, got %v instead", c1, c2, expected, c1)
	}
}

func TestComplexMul(t *testing.T) {
	var c1 = Complex{1, 1}
	var c2 = Complex{2, 2}
	var c3 = Mul(&c1, &c2)

	var expected = complex(1, 1) * complex(2, 2)

	if almostEqual(float64(c3.im), imag(expected)) && almostEqual(float64(c3.re), real(expected)) {
		t.Errorf("%v * %v should be %v, got %v instead", c1, c2, expected, c3)
	}
}

func TestComplexMulMethod(t *testing.T) {
	var c1 = Complex{1, 1}
	var c2 = Complex{2, 2}
	c1.Mul(&c2)

	var expected = complex(1, 1) * complex(2, 2)

	if almostEqual(float64(c1.im), imag(expected)) && almostEqual(float64(c1.re), real(expected)) {
		t.Errorf("%v * %v should be %v, got %v instead", c1, c2, expected, c1)
	}
}

func TestComplexDiv(t *testing.T) {
	var c1 = Complex{1, 1}
	var c2 = Complex{2, 2}
	var c3, _ = Div(&c1, &c2)

	var expected = complex(1, 1) / complex(2, 2)

	if almostEqual(float64(c3.im), imag(expected)) && almostEqual(float64(c3.re), real(expected)) {
		t.Errorf("%v * %v should be %v, got %v instead", c1, c2, expected, c3)
	}
}

func TestComplexDivMethod(t *testing.T) {
	var c1 = Complex{1, 1}
	var c2 = Complex{2, 2}
	_ = c1.Div(&c2)

	var expected = complex(1, 1) / complex(2, 2)

	if almostEqual(float64(c1.im), imag(expected)) && almostEqual(float64(c1.re), real(expected)) {
		t.Errorf("%v * %v should be %v, got %v instead", c1, c2, expected, c1)
	}
}
