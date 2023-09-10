package generics

import "fmt"

func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

type number interface {
    int64 | float64
}

func SumIntsOrFloats[K comparable, V number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

func RunGenericsDemo() {
    m1 := map[string]int64 {
        "a": 1,
        "b": 2,
    }
     
    m2 := map[string]float64 {
        "a": 1.0,
        "b": 2.2,
    }

    fmt.Println(SumInts(m1), "\n===========\n", SumFloats(m2))

    fmt.Println("\nentering Generics...")

    fmt.Println(
        SumIntsOrFloats[string, int64](m1),
        SumIntsOrFloats[string, float64](m2),
    )
}
