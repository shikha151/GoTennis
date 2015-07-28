package main
import "fmt"
import "os"

type ScoreBoard struct {
   playerAScore int
   playerBScore int
   setsWonByA int
   setsWonByB int
}
var scoreBoard ScoreBoard

func main() {
scoreBoard.playerAScore=0
scoreBoard.playerBScore=0
scoreBoard.setsWonByA=0
scoreBoard.setsWonByB=0	
	
getScore("ABABABBBBAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
}

func getScore(scorePattern string){
for i := 0; i < len(scorePattern); i++ {
      updateScore(scorePattern[i])
   }
}

func updateScore(score byte){
   if (score=='A') {
      scoreBoard.playerAScore++
   } 
   if (score=='B'){
      scoreBoard.playerBScore++
   }
   if (hasWonSet()) {
   	scoreBoard.playerAScore=0
   	scoreBoard.playerBScore=0
   } else {
   printScores()
   }
}


func printScores() {
	score := []string{"0", "15", "30", "40", "A"}

		if ((scoreBoard.playerAScore < 4) && (scoreBoard.playerBScore < 4)) {
			fmt.Println(score[scoreBoard.playerAScore] + "-" + score[scoreBoard.playerBScore])
		}else {
			if (scoreBoard.playerAScore == scoreBoard.playerBScore) {
				fmt.Println("40-40")
			}else if (scoreBoard.playerAScore > scoreBoard.playerAScore) {
			fmt.Println("A-40")
			}else {
			fmt.Println("40-A")
			}
		}
}

func hasWonSet() bool{
   
   var aScore int = scoreBoard.playerAScore;
   var bScore int = scoreBoard.playerBScore;
 
   if( (aScore > 3) && (aScore - bScore >= 2) ) {
       scoreBoard.setsWonByA++
       fmt.Printf("%d-%d\n",scoreBoard.setsWonByA,scoreBoard.setsWonByB );
	if(!hasWonMatch()) {
	   return true
        } else {
         fmt.Printf("Game Over\n");  
         os.Exit(1)
         }
   } else if( (bScore > 3) && (bScore - aScore >= 2) ){
       scoreBoard.setsWonByB++
       fmt.Printf("%d-%d\n",scoreBoard.setsWonByA,scoreBoard.setsWonByB );
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
    
    var aSetsWon int = scoreBoard.setsWonByA;
    var bSetsWon int = scoreBoard.setsWonByB;
    
    if( (aSetsWon > 5) && (aSetsWon - bSetsWon >= 2) ) {
       fmt.Printf("Player A wins\n" );
	   return true
   } else if( (bSetsWon > 5) && (bSetsWon - aSetsWon >= 2) ){
       fmt.Printf("Player B Wins\n" );
	   return true
   }
   return false
}
