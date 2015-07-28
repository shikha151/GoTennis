package main
import "fmt"

type Game struct {
   playerAScore int
   playerBScore int
}
var myGame Game


func main() {
myGame.playerAScore=0
myGame.playerBScore=0

updateScore("A")
}



func updateScore(score string){
   if (score=="A") {
      game.playerAScore++
   } 
   if (score=="B"){
      game.playerBScore++
   }   
   fmt.Println("Score Updated")
}


func printScores() {
	score := []strings{"0", "15", "30", "40", "A"}
	isWinner := false
	if (isWinner==false) {

		if ((myGame.playerAScore < 4) && (myGame.playerBScore < 4)) {
			fmt.Println(score[myGame.playerAScore] + "-" + score[myGame.playerBScore])
		}
		else {
			if (myGame.playerAScore == myGame.playerBScore) {
				fmt.Println("40-40")
			}
			else if (myGame.playerAScore > myGame.playerAScore) {
			fmt.Println("A-40")
			}
			else {
			fmt.Println("40-A")
			}
		}
	}
}
