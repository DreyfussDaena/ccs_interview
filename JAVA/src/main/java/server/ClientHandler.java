package server;

import server.GameLogic;
import java.io.*;
import java.net.*;

class ClientHandler extends Thread {
    private final Socket clientSocket;
    private final GameLogic gameLogic;

    public ClientHandler(Socket socket) {
        this.clientSocket = socket;
        this.gameLogic = new GameLogic();
    }

    public void run() {
        try (BufferedReader in = new BufferedReader(new InputStreamReader(clientSocket.getInputStream()));
             PrintWriter out = new PrintWriter(clientSocket.getOutputStream(), true)) {

            out.println("Welcome to the Guessing Game! Enter a number between 1 and 100.");
            String inputLine;

            while ((inputLine = in.readLine()) != null) {
                try {
                    int guess = gameLogic.validateGuess(inputLine);
                    boolean isCorrect = gameLogic.checkGuessCorrectness(guess);
                    String prefix = "";
//                     String prefix = gameLogic.generatePrefix(guess);

                    if (isCorrect) {
                        out.println(prefix + " Congratulations! You guessed correctly!");
                        break;
                    } else {
                        out.println(prefix + " Try again!");
                    }
                } catch (IllegalArgumentException e) {
                    out.println(e.getMessage());
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            try {
                clientSocket.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
}
