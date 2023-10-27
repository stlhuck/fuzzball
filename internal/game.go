package game

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	wr "github.com/mroth/weightedrand"
)

var pause = 3 * time.Second

type Game struct {
	HalfInning  int
	Inning      int
	Outs        int
	AwayHits    int
	AwayScore   int
	AwayErrors  int
	HomeHits    int
	HomeScore   int
	HomeErrors  int
	TopOrBottom string
}

func Clear() {
	screen.Clear()
	screen.MoveTopLeft()
}

func PlayBall() {
	g := Game{
		HalfInning:  1,
		Inning:      1,
		Outs:        0,
		AwayHits:    0,
		AwayScore:   0,
		AwayErrors:  0,
		HomeHits:    0,
		HomeScore:   0,
		HomeErrors:  0,
		TopOrBottom: "",
	}

	fmt.Printf("Here we go as we lead things off in the top of the 1st")

	for g.Inning < 7 {
		HalfInning(&g)

		switch g.HalfInning {
		case 0:
			g.TopOrBottom = "top"
		case 1:
			g.TopOrBottom = "bottom"

		}

		fmt.Printf("Going to the %v of the %v\n", g.TopOrBottom, g.Inning/2)
		g.Outs = 0
	}
	fmt.Println("That's Game")
}

func HalfInning(g *Game) {
	Clear()
	// Initializing a game
	// This is where we will keep track of innings, score, hits, and errors as we progress.

	// This is the tracking of a plate apperance

	/*halfInning:  Should track everything that happens within 3 outs
	This such as number of outs, hits, runs, errors, etc.
	Currently just working on tracking the number of outs*/

	for g.Outs < 3 {
		atBat(g)
		fmt.Printf("Inning: %v\n", g.Inning/2)
		fmt.Printf("Outs: %v\n", g.Outs)

	}
	g.Inning++
}

func atBat(g *Game) {
	var walk, strikeout int

	chooser, _ := wr.NewChooser(
		wr.NewChoice("Ball", 9),
		wr.NewChoice("Strike", 9),
		wr.NewChoice("Ball In Play", 3),
	)

	for p := 1; p <= 30; p++ {

		pitch := chooser.Pick()

		switch {
		case pitch == "Ball":
			walk++
			fmt.Printf("Pitch %v:\tBall %v\n", p, walk)

			if walk == 4 {
				fmt.Println("That's a walk, take your base.")
				fmt.Println()
				return
			}
		case pitch == "Strike":
			strikeout++
			fmt.Printf("Pitch %v:\tStrike %v\n", p, strikeout)

			if strikeout == 3 {
				fmt.Println("Strike 3, You're out!")
				g.Outs++
				fmt.Println()
				if g.Outs == 3 {
					g.Inning++
					fmt.Println(g.Outs, g.Inning)
					fmt.Println()
					fmt.Println()
				}
				return
			}
		case pitch == "Ball In Play":
			fmt.Printf("Pitch %v:\tThe ball is in play\n", p)
			fmt.Println()
			return

		}
		time.Sleep(pause)

	}

}
