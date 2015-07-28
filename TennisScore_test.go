package main
import "testing"

var myGameTest Game

func TestUpdateScore(t *testing.T) {
	myGameTest.playerAScore=0;
	updateScore('A')
	if (myGameTest.playerAScore==1) {

		t.Error("Expected 1 got ", myGameTest.playerAScore)

	}

}

func TestHasWonSet(t *testing.T) {
	myGameTest.playerAScore=0;
	myGameTest.playerBScore=0;
	getScore("ABBA")
	if (hasWonSet()) {

		t.Error("Expected False got ", hasWonSet())

	}

}

func TestHasWonMatch(t *testing.T) {
	myGameTest.playerAScore=0;
	myGameTest.playerBScore=0;
	getScore("AAAAAAAAAAAAAAAAAAAAAAAA")
	if (!hasWonSet()) {

		t.Error("Expected True got ", hasWonSet())

	}

}

func TestHasWonSet2(t *testing.T) {
	myGameTest.playerAScore=0;
	myGameTest.playerBScore=0;
	getScore("ABBA")
	if (hasWonSet()) {

		t.Error("Expected False got ", hasWonSet())

	}

}
