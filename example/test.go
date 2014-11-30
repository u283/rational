package main
import (
	"fmt"
	"rational"
)

func main() {
	var a rational.Rational
	var b rational.Rational

	fmt.Print("Enter a: ")
	fmt.Scanf("%d/%d", &a.Num, &a.Den)
	fmt.Print("Enter b: ")
	fmt.Scanf("%d/%d", &b.Num, &b.Den)

	var c rational.Rational

	c = rational.Add(a, b)
	fmt.Printf("a + b = %d / %d\n", c.Num, c.Den)

	c = rational.Subtract(a, b)
	fmt.Printf("a - b = %d / %d\n", c.Num, c.Den)

	c = rational.Multiply(a, b)
	fmt.Printf("a * b = %d / %d\n", c.Num, c.Den)

	c = rational.Divide(a, b)
	fmt.Printf("a / b = %d / %d\n", c.Num, c.Den)
}
