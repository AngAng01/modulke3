package main

import (
    "fmt"
)

func factorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}

func permutation(a, c int) int {
    return factorial(a) / factorial(a-c)
}

func combination(a, c int) int {
    return factorial(a) / (factorial(c) * factorial(a-c))
}

func main() {
    var a, b, c, d int

    fmt.Println("Masukkan 4 angka (a, b, c, d): ")
    fmt.Scan(&a, &b, &c, &d)

    permAC := permutation(a, c)
    combAC := combination(a, c)

    permBD := permutation(b, d)
    combBD := combination(b, d)

    fmt.Printf("P(%d, %d) = %d\n", a, c, permAC)
    fmt.Printf("C(%d, %d) = %d\n", a, c, combAC)
    fmt.Printf("P(%d, %d) = %d\n", b, d, permBD)
    fmt.Printf("C(%d, %d) = %d\n", b, d, combBD)
}
