package main
import "fmt"

type Game struct {
   playerAScore int
   playerBScore int
   setsWonByA int
   setsWonByB int
}
var myGame Game

func main() {
myGame.playerAScore=0
myGame.playerBScore=0
myGame.setsWonByA=0
myGame.setsWonByB=0

var scorePattern string
fmt.Scanln(&scorePattern)
fmt.Println("\n")
getScore(scorePattern)
updateScore('B')
}

func getScore(scorePattern string){
for i := 0; i < len(scorePattern); i++ {
      updateScore(scorePattern[i])
   }
}

func updateScore(score byte){
   if (score=='A') {
      myGame.playerAScore++
   } 
   if (score=='B'){
      myGame.playerBScore++
   }
   if (hasWonSet()) {
   	myGame.playerAScore=0
   	myGame.playerBScore=0
   } else {
   printScores()
   }
}


func printScores() {
	score := []string{"0", "15", "30", "40", "A"}

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

func hasWonSet() bool{
   
   var aScore int = myGame.playerAScore;
   var bScore int = myGame.playerBScore;
 
   if( (aScore > 3) && (aScore - bScore >= 2) ) {
       myGame.setsWonByA++
       fmt.Printf("%d-%d\n",myGame.setsWonByA,myGame.setsWonByB );
	   return true
   } else if( (bScore > 3) && (bScore - aScore >= 2) ){
       myGame.setsWonByB++
       fmt.Printf("%d-%d\n",myGame.setsWonByA,myGame.setsWonByB );
	   return true
   }
   return false
}
