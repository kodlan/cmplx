package cmplx

import "errors"

type Complex struct {
	re float32
	im float32
}

func Add(c1, c2 *Complex) Complex {
	return Complex{re: c1.re + c2.re, im: c1.im + c2.im}
}

func (c1 *Complex) Add(c2 *Complex) {
	c1.re = c1.re + c2.re
	c1.im = c1.im + c2.im
}

func Sub(c1, c2 *Complex) Complex {
	return Complex{re: c1.re - c2.re, im: c1.im - c2.im}
}

func (c1 *Complex) Sub(c2 *Complex) {
	c1.re = c1.re - c2.re
	c1.im = c1.im - c2.im
}

func Mul(c1, c2 *Complex) Complex {
	return Complex{
		re: c1.re*c2.re - c1.im*c2.im,
		im: c1.re*c2.im + c1.im*c2.re}
}

func (c1 *Complex) Mul(c2 *Complex) {
	c1.re = c1.re*c2.re - c1.im*c2.im
	c1.im = c1.re*c2.im + c1.im*c2.re
}

func Div(c1, c2 *Complex) (Complex, error) {
	divider := c2.im*c2.im + c2.re*c2.re

	if divider == 0 {
		return Complex{}, errors.New("can't divide by 0")
	}

	return Complex{
		re: (c1.re*c2.re + c1.im*c2.im) / divider,
		im: (c1.im*c2.re - c1.re*c2.im) / divider}, nil
}

func (c1 *Complex) Div(c2 *Complex) error {
	divider := c2.im*c2.im + c2.re*c2.re

	if divider == 0 {
		return errors.New("can't divide by 0")
	}

	c1.re = (c1.re*c2.re + c1.im*c2.im) / divider
	c1.im = (c1.im*c2.re - c1.re*c2.im) / divider

	return nil
}
