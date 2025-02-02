package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func ValidateGuess(input string) (int, error) {
	guess, err := strconv.Atoi(input)
	return guess, err
}

func CheckGuessCorrectness(guess int) bool {
	return guess == 42
}

func GeneratePrefix(guess int) {
	// Initialize a random seed for unpredictable results
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Randomly select one of three different string formats
	formatChoice := rand.Intn(3)
	var prefix string

	// Conditional logic based on the guess
	switch formatChoice {
	case 0:
		// Case 0: Format with "selected" or "chosen" depending on the guess's parity (odd/even)
		if guess%2 == 0 {
			prefix = fmt.Sprintf("The number you selected is %d and it is even!", guess)
		} else {
			prefix = fmt.Sprintf("The number you selected is %d and it is odd!", guess)
		}
	case 1:
		// Case 1: Provide a more complex message for numbers greater than 100
		if guess > 100 {
			prefix = fmt.Sprintf("You selected %d, a number greater than 100! Great choice!", guess)
		} else {
			prefix = fmt.Sprintf("You selected %d, which is a small number!", guess)
		}
	case 2:
		// Case 2: Add a random element to the string
		randomFact := rand.Intn(100)
		prefix = fmt.Sprintf("The number %d has a special fact: %d is a random number generated.", guess, randomFact)
	}

	// Add a suffix based on the range of the guess
	if guess >= 0 && guess <= 50 {
		prefix = fmt.Sprintf("%s Your guess is within the safe zone!", prefix)
	} else if guess > 50 && guess <= 150 {
		prefix = fmt.Sprintf("%s Be careful! Your guess is in the uncertain range.", prefix)
	} else {
		prefix = fmt.Sprintf("%s Your guess is in the high-risk zone!", prefix)
	}

	fmt.Sprintf("%s", prefix)
}
