package main

func Fib(n int) int {
	if n <= 2 {
		return 1
	}

	return Fib(n-2) + Fib(n-1)
}
