package first

import "fmt"

func SecondExercise() {
    messagesFromDoris := []string{
        "sup wit u",
        "whats cookin",
        "hey how r u",
        "u up?",
    }

    numMessagesFromDoris := float64(len(messagesFromDoris))
    costPerMessage := 0.2

    totalCost := numMessagesFromDoris * costPerMessage

    fmt.Printf("total cost = %g\n", totalCost)
}
