package main
import "fmt"

type Game struct {
   playerAScore int
   playerBScore int
}
var game Game


func main() {
game.playerAScore=0
game.playerBScore=0

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
