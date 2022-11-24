package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Player struct {
	FirstName string
	LastName  string
	Pseudo    string
	Score     int
}

func main() {
	if _, err := os.Stat("score.txt"); err == nil {
		showOldScore()
	}
	startGame()
}

func startGame() {
	players := savePlayer()
	nbRounds := 5
	for i := 0; i < nbRounds; i++ {
		if i == 0 {
			clearConsole()
		}
		fmt.Printf("\n-------ROUND %d -------\n", i+1)
		playGame(players)
	}
	saveScore(players)
}

func savePlayer() []Player {
	players := []Player{}
	player := Player{}
	temp := ""
	for temp != "fin" {
		fmt.Printf("\nEntrez le prénom : ")
		fmt.Scan(&temp)
		player.FirstName = temp

		fmt.Printf("\nEntrez le nom : ")
		fmt.Scan(&temp)
		player.LastName = temp

		fmt.Printf("\nEntrez le pseudo : ")
		fmt.Scan(&temp)
		player.Pseudo = temp

		players = append(players, player)

		fmt.Printf("\nIl y a %d joueurs.\n", len(players))

		for i := 0; i < len(players); i++ {
			fmt.Printf("Le joueur %d s'appelle %s %s, il porte le pseudo %s et possède un score de %d.\n", i+1, players[i].FirstName, players[i].LastName, players[i].Pseudo, players[i].Score)
		}

		fmt.Printf("\nEntrez 'fin' pour terminer la saisie des joueurs, ou entrez c pour continuer : ")
		fmt.Scan(&temp)

	}

	return players
}

func playGame(players []Player) {
	if len(players) >= 1 {
		rand.Seed(time.Now().UnixNano())
		rnumber := rand.Intn(100) + 1
		tempnumber := -1

		for tempnumber != rnumber {
			for i := 0; i < len(players); i++ {
				fmt.Printf("\n%s, c'est à toi de jouer ! (%d points)\n", players[i].Pseudo, players[i].Score)
				fmt.Printf("Entrez un nombre entre 1 et 100 : ")
				fmt.Scan(&tempnumber)
				if tempnumber == rnumber {
					fmt.Printf("Bravo %s, tu as trouvé le nombre mystère !\n", players[i].Pseudo)
					players[i].Score += 10
					showScore(players)
					break
				} else if tempnumber > rnumber {
					fmt.Printf("Le nombre mystère est plus petit que %d.\n", tempnumber)
					players[i].Score -= 1
				} else {
					fmt.Printf("Le nombre mystère est plus grand que %d.\n", tempnumber)
					players[i].Score -= 1
				}
			}
		}
	}
}

func showScore(players []Player) {
	fmt.Printf("\n------- SCORES -------\n\n")
	for i := 0; i < len(players); i++ {
		fmt.Printf("%s : %d\n", players[i].Pseudo, players[i].Score)
	}
	fmt.Printf("\n------- SCORES -------\n")
}

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

func saveScore(Players []Player) {
	file, err := os.Create("score.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < len(Players); i++ {
		fmt.Fprintf(file, "%s : %d\n", Players[i].Pseudo, Players[i].Score)
	}
}

func showOldScore() {

	fmt.Printf("\n------- HIGHT SCORES -------\n\n")

	file, err := os.Open("score.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		fmt.Print(string(buf[:n]))
	}

	fmt.Printf("\n----------------------------\n\n")
}
