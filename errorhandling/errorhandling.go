package errorhandling

import "fmt"

type number interface {
    int | float64
}

type divideError struct {
    divisor interface{}
    dividend interface{}
}

func (d divideError) Error() string {
    return fmt.Sprintf(
        "%v can't be divided by %v",
        d.dividend,
        d.divisor,
    )
}

func divide[n number](dividend, divisor n) (n, error) {
   if divisor == 0 {
       return 0, divideError{
        divisor: divisor,
        dividend: dividend,
        }
    }
    return dividend/divisor, nil
}

func RunErrorHandlingDemo() {
    fmt.Println(divide(22.4, 45))
    fmt.Println(divide(22, 13))
    fmt.Println(divide(22, 11))
    fmt.Println(divide(13, 0))
}
