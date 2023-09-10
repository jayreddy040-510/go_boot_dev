package errorhandling

import "fmt"

type number interface {
    int | float32
}

type divideError struct {
    divisor number
    dividend number
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
    
}
