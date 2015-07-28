package main
import "fmt"
import "os"

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
getScore("ABABABBBBAAAA")
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
	if(!hasWonMatch()) {
	   return true
        } else {
         fmt.Printf("Game Over\n");  
         os.Exit(1)
         }
   } else if( (bScore > 3) && (bScore - aScore >= 2) ){
       myGame.setsWonByB++
       fmt.Printf("%d-%d\n",myGame.setsWonByA,myGame.setsWonByB );
       if(!hasWonMatch()) {
	   return true
       } else {
         fmt.Printf("Game Over\n");
         os.Exit(1)
         }
   }
   return false
}

func hasWonMatch() bool{
    
    var aSetsWon int = myGame.setsWonByA;
    var bSetsWon int = myGame.setsWonByB;
    
    if( (aSetsWon > 5) && (aSetsWon - bSetsWon >= 2) ) {
       fmt.Printf("Player A wins\n" );
	   return true
   } else if( (bSetsWon > 5) && (bSetsWon - aSetsWon >= 2) ){
       fmt.Printf("Player B Wins\n" );
	   return true
   }
   return false
}
