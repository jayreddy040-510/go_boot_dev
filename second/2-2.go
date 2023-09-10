package second

import "fmt"

type msgToSend struct {
    phoneNumber int
    message string
}

func test(m msgToSend) {
    fmt.Printf(
        "Sending message: '%s' to: %d\n",
        m.message,
        m.phoneNumber,
    )
    fmt.Println("========================")
}
func SecondExercise() {
   test(msgToSend{
        phoneNumber: 5598857595,
        message: "sup dude",
    }) 
}
