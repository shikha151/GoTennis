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
      myGame.playerAScore++
   } 
   if (score=="B"){
      myGame.playerBScore++
   }   
   fmt.Println("Score Updated")
	printScores()
}


func printScores() {
	score := []string{"0", "15", "30", "40", "A"}
	if (!isWinner()) {

		if ((myGame.playerAScore < 4) && (myGame.playerBScore < 4)) {
			fmt.Println(score[myGame.playerAScore] + "-" + score[myGame.playerBScore])
		}else {
			if (myGame.playerAScore == myGame.playerBScore) {
				fmt.Println("40-40")
			}else if (myGame.playerAScore > myGame.playerAScore) {
			fmt.Println("A-40")
			}else {
			fmt.Println("40-A")
			}
		}
	}
}

func isWinner() bool{
   
   var aScore int = myGame.playerAScore;
   var bScore int = myGame.playerBScore;
 
   if( (aScore > 3) && (aScore - bScore >= 2) ) {
       fmt.Printf("1-0\n" );
	   return true
   } else if( (bScore > 3) && (bScore - aScore >= 2) ){
       fmt.Printf("0-1\n" );
	   return true
   }
   return false
}
