package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var n int = 1

	fmt.Print("Fibonacci number: ")
	fmt.Scanf("%d\n", &n)

	var f int
	var start time.Time
	var duration time.Duration

	fmt.Println("Running recursive Fibonacci ...")
	start = time.Now()
	f = recursiveFibonacci(n)
	duration = time.Since(start)
	fmt.Printf("%dth Fibonacci number is %d, it took %d ns to compute\n\n", n, f, duration.Nanoseconds())

	fmt.Println("Running dynamic Fibonacci ...")
	start = time.Now()
	f = dynamicFibonacci(n)
	duration = time.Since(start)
	fmt.Printf("%dth Fibonacci number is %d, it took %d ns to compute\n\n", n, f, duration.Nanoseconds())

	fmt.Println("Running binet Fibonacci ...")
	start = time.Now()
	f = binetFibonacci(n)
	duration = time.Since(start)
	fmt.Printf("%dth Fibonacci number is %d, it took %d ns to compute\n\n", n, f, duration.Nanoseconds())
}

func recursiveFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
}

func dynamicFibonacci(n int) int {
	arr := make([]int, n+1)

	arr[0] = 0
	arr[1] = 1

	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}

func binetFibonacci(n int) int {
	a := (1 + math.Sqrt(5)) / 2
	b := (1 - math.Sqrt(5)) / 2
	f := (math.Pow(a, float64(n)) - math.Pow(b, float64(n))) / (a - b)

	return int(math.Round(f))
}
