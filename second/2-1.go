package second

import "fmt"

func FirstExercise() {

   var smsSendingLimit uint8
   var costPerSMS float32
   var hasPermission bool
   var username string

    fmt.Printf(
        "%v %.2f %v %q\n",
        smsSendingLimit,
        costPerSMS,
        hasPermission,
        username,
    )
}
