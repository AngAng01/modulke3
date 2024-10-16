package main

import (
    "fmt"
    "math"
)

func distance(x1, y1, x2, y2 int) float64 {
    return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func main() {
    var cx1, cy1, r1 int
    fmt.Scan(&cx1, &cy1, &r1)

    var cx2, cy2, r2 int
    fmt.Scan(&cx2, &cy2, &r2)

    var x, y int
    fmt.Scan(&x, &y)

    d1 := distance(x, y, cx1, cy1)
    d2 := distance(x, y, cx2, cy2)

    inCircle1 := d1 <= float64(r1)
    inCircle2 := d2 <= float64(r2)

    if inCircle1 && inCircle2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if inCircle1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if inCircle2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}
