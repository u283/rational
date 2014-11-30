package rational

type Rational struct {
	Num int
	Den int
}

func Reduce(a Rational) Rational {
	var min int
	if a.Num < a.Den {
		min = a.Num
	} else {
		min = a.Den
	}
	
	var maxdivider int
	for i := min; i > 1; i-- {
		if ((a.Num % i) == 0) && ((a.Den % i) == 0) {
			maxdivider = i
			break
		}
	}
	
	var b Rational
	b.Num = a.Num / maxdivider
	b.Den = a.Den / maxdivider
	return b
}

func Multiply(a, b Rational) Rational {
	var c Rational
	c.Num = a.Num * b.Num
	c.Den = a.Den * b.Den
	return Reduce(c)
}

func Divide(a, b Rational) Rational {
	var c Rational
	c.Num = a.Num * b.Den
	c.Den = a.Den * b.Num
	return Reduce(c)
}

func Add(a, b Rational) Rational {
	var c Rational
	c.Num = (a.Num * b.Den) + (b.Num * a.Den)
	c.Den = a.Den * b.Den
	return Reduce(c)
}

func Subtract(a, b Rational) Rational {
	var c Rational
	c.Num = (a.Num * b.Den) - (b.Num * a.Den)
	c.Den = a.Den * b.Den
	return Reduce(c)
}
