package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	var n int = 5

	fmt.Print("Primes limit: ")
	fmt.Scanf("%d\n", &n)

	var p []bool
	var start time.Time
	var duration time.Duration

	fmt.Println("Running Algorithm 1 ...")
	start = time.Now()
	p = alg1(n)
	duration = time.Since(start)
	fmt.Printf("Algorithm 1 took %d ns to compute for the limit of %d\n", duration.Nanoseconds(), n)
	printPrimes(p)

	fmt.Println("Running Algorithm 2 ...")
	start = time.Now()
	p = alg2(n)
	duration = time.Since(start)
	fmt.Printf("Algorithm 2 took %d ns to compute for the limit of %d\n", duration.Nanoseconds(), n)
	printPrimes(p)

	fmt.Println("Running Algorithm 3 ...")
	start = time.Now()
	p = alg3(n)
	duration = time.Since(start)
	fmt.Printf("Algorithm 3 took %d ns to compute for the limit of %d\n", duration.Nanoseconds(), n)
	printPrimes(p)

	fmt.Println("Running Algorithm 4 ...")
	start = time.Now()
	p = alg4(n)
	duration = time.Since(start)
	fmt.Printf("Algorithm 4 took %d ns to compute for the limit of %d\n", duration.Nanoseconds(), n)
	printPrimes(p)

	fmt.Println("Running Algorithm 5 ...")
	start = time.Now()
	p = alg5(n)
	duration = time.Since(start)
	fmt.Printf("Algorithm 5 took %d ns to compute for the limit of %d\n", duration.Nanoseconds(), n)
	printPrimes(p)
}

func printPrimes(c []bool) {
	for i := 0; i < len(c); i++ {
		if !c[i] {
			fmt.Print(i, " ")
		}
	}
	fmt.Print("\n\n")
}

func alg1(n int) []bool {
	c := make([]bool, n+1)
	c[1] = true

	for i := 2; i <= n; i++ {
		if c[i] == false {
			for j := 2 * i; j <= n; j = j + i {
				c[j] = true
			}
		}
	}

	return c
}

func alg2(n int) []bool {
	c := make([]bool, n+1)
	c[1] = true

	for i := 2; i <= n; i++ {
		for j := 2 * i; j <= n; j = j + i {
			c[j] = true
		}
	}

	return c
}

func alg3(n int) []bool {
	c := make([]bool, n+1)
	c[1] = true

	for i := 2; i <= n; i++ {
		if c[i] == false {
			for j := i + 1; j <= n; j++ {
				if j%i == 0 {
					c[j] = true
				}
			}
		}
	}

	return c
}

func alg4(n int) []bool {
	c := make([]bool, n+1)
	c[1] = true

	for i := 2; i <= n; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				c[i] = true
			}
		}
	}

	return c
}

func alg5(n int) []bool {
	c := make([]bool, n+1)
	c[1] = true

	for i := 2; i <= n; i++ {
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				c[i] = true
			}
		}
	}

	return c
}
