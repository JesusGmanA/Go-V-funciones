package main

import "fmt"

func sum(args ...int) int {
	max := 0
	for _, val := range args {
		if max < val {
			max = val
		}
	}
	return max

}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func intercambia(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func generadorImpares() func() uint {
	i := uint(1)
	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

func main() {
	x := []int{15, 1, 4, 2, 22, 12}
	a := 15
	b := 10
	fmt.Println("Valor max del slice: ", sum(x...))
	fmt.Println("Fibonacci: ", Fibonacci(20))
	intercambia(&a, &b)
	fmt.Println("Valor de a: ", a)
	fmt.Println("Valor de b: ", b)
	nextImpar := generadorImpares()
	fmt.Println("Impar: ", nextImpar())
	fmt.Println("Impar: ", nextImpar())
	fmt.Println("Impar: ", nextImpar())
	fmt.Println("Impar: ", nextImpar())
}
