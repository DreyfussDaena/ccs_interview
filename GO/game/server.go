package game

import (
	"fmt"
	"log"
	"net"
)

func StartServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server started, waiting for a player...")

	// Accept only one player connection for now.
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Error accepting connection: %v", err)
	}
	defer conn.Close()

	fmt.Println("Player has connected.")

	// Simple game logic to check the player's guess
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Error reading from client: %v", err)
			return
		}

		// Process the guess sent by the client (assuming it's a number)
		guess := string(buffer[:n])
		fmt.Printf("Received guess: %s\n", guess)

		numGuess, err := ValidateGuess(guess)
		if err != nil {
			log.Printf("Error validating guess: %v", err)
			writeToClient(conn, err.Error())
		} else {
			// Check if the guess matches the correct answer
			var response, prefix string
			if ValidateGuessCorrectness(numGuess) {
				// prefix = GeneratePrefix()
				response = "Congratulations! You guessed the correct number!"
			} else {
				response = "Try again!"
			}

			// Send the response back to the client
			writeToClient(conn, prefix+response)
		}
	}
}

func writeToClient(conn net.Conn, s string) {
	_, err := conn.Write([]byte(s))
	if err != nil {
		log.Printf("Error writing to client: %v", err)
		return
	}
}
