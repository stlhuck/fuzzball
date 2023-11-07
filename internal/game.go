package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/inancgumus/screen"
	wr "github.com/mroth/weightedrand"
)

var pause = 2 * time.Second

type Game struct {
	HalfInning int
	Inning     int
	InningDesc string
	Outs       int
	AwayHits   int
	AwayScore  int
	AwayErrors int
	HomeHits   int
	HomeScore  int
	HomeErrors int
}

//  actually maybe I make these bools and let's try that
type Bases struct {
	FirstBase  bool
	SecondBase bool
	ThirdBase  bool
}

func Clear() {
	screen.Clear()
	screen.MoveTopLeft()
}

func PlayBall() {
	g := Game{
		HalfInning: 1,
		Inning:     1,
		InningDesc: "",
		Outs:       0,
		AwayHits:   0,
		AwayScore:  0,
		AwayErrors: 0,
		HomeHits:   0,
		HomeScore:  0,
		HomeErrors: 0,
	}

	for g.Inning < 3 {
		//Clear()
		FullInning(&g)

		// Increasing the inning value by one as we move forward through the game
		g.Inning++

		// Resettig the half-inning variable for tracking purposes
		g.HalfInning = 1
	}

	fmt.Println("That's Game")
}

// Tracking which bases are occupied and their behavior during an a play
// This includes when a runner scores, we add to their teams total score
// In the future, we will track who is on which base (maybe ghost runners, haven't decided)
// The logic starts are thirdbase and works backwards so that all bases are properly assigned runners
func BaseTracking(g *Game, b *Bases, numOfBases int) {

	// Handling a runner on thirdbase
	if b.ThirdBase && g.InningDesc == "top" {
		g.AwayScore++
		b.ThirdBase = false
	} else if b.ThirdBase && g.InningDesc == "bottom" {
		g.HomeScore++
		b.ThirdBase = false
	}

	// Handling a runner on second base
	if b.SecondBase {
		if numOfBases >= 2 {
			if g.InningDesc == "top" {
				g.AwayScore++
				b.SecondBase = false
			} else if g.InningDesc == "bottom" {
				g.HomeScore++
				b.SecondBase = false
			}
		} else if numOfBases == 1 {
			b.ThirdBase = true
			b.SecondBase = false
		}
	}

	// Handling a runner on first base
	if b.FirstBase {
		if numOfBases >= 3 {
			if g.InningDesc == "top" {
				g.AwayScore++
				b.FirstBase = false
			} else if g.InningDesc == "bottom" {
				g.HomeScore++
				b.FirstBase = false
			}
		} else if numOfBases == 2 {
			b.ThirdBase = true
			b.FirstBase = false
		} else if numOfBases == 1 {
			b.SecondBase = true
			b.FirstBase = false
		}
	}

	// Handling no runner on first base {
	if !b.FirstBase {
		if numOfBases == 4 {
			if g.InningDesc == "top" {
				g.AwayScore++
				b.SecondBase = false
			} else if g.InningDesc == "bottom" {
				g.HomeScore++
				b.SecondBase = false
			}
		} else if numOfBases == 3 {
			b.ThirdBase = true
		} else if numOfBases == 2 {
			b.SecondBase = true
		} else if numOfBases == 1 {
			b.FirstBase = true
		}
	}
}

func FullInning(g *Game) {
	// The Full Inning fuction, will contain information provided from 2 HalfInnings (6 outs)
	// [ ] # of runs for away team - Phase 2
	// [ ] # of hits for away team - Phase 2
	// [ ] # of errors for away team - Phase 2
	// [ ] # of runs for home team - Phase 2
	// [ ] # of hits for home team - Phase 2
	// [ ] # of errors for home team - Phase 2
	//

	for g.HalfInning <= 2 {
		g.Outs = 0
		HalfInning(g)
	}

}

func HalfInning(g *Game) {
	// The halfInning fuction, will contain information related to the top or bottom of an inning (3 outs)
	// [x] # of outs - MVP
	// [x] Identify if we are in the top or bottom of an inning - MVP
	// [] # of hits - Phase 2
	// [] # of runs - Phase 2
	// [] # of errors - Phase 3+

	// Determining if we are in the top half or bottom half of an inning.
	topOrBottom := g.HalfInning % 2

	if topOrBottom == 0 {
		g.InningDesc = "bottom"
	} else {
		g.InningDesc = "top"
	}

	// Printing out a summary for between half innings  TODO: Move this to it's own func and make it a runs, hits, errors display, also display it at the end of the game.
	fmt.Println("*****\tGame Summary\t*****")
	fmt.Printf("Going into the %v of the %v\n", g.InningDesc, g.Inning)
	fmt.Printf("Away Score:\t%v\n", g.AwayScore)
	fmt.Printf("HomeScore:\t%v\n", g.HomeScore)
	fmt.Println("******************************")
	fmt.Println()

	//resetting the bases TODO: Leaving off here as I was try to fiugre out how to manage tracking the bases.  Maybe a fresh start will help.
	b := Bases{
		FirstBase:  false,
		SecondBase: false,
		ThirdBase:  false,
	}

	for g.Outs < 3 {
		atBat(g, &b)
		fmt.Println()
		// fmt.Printf("Debug Summary\ng.Inning:\t%v\ng.Outs:\t%v\ng.HalfInnning:\t%v\ng.InningDesc:\t%v\nb.Bases:\t%v\ng.AwayScore:\t%v\ng.HomeScore:\t%v\n", g.Inning, g.Outs, g.HalfInning, g.InningDesc, b, g.AwayScore, g.HomeScore)
		// fmt.Println()
		// time.Sleep(pause)
		// time.Sleep(pause)
		// time.Sleep(pause)
	}
}

// The atBat function allows for the action between a pitcher and batter
// The action can result in hits, runs, outs, etc.
func atBat(g *Game, b *Bases) {
	//time.Sleep(pause)
	var walk, strikeout int

	chooser, _ := wr.NewChooser(
		wr.NewChoice("Ball", 2),
		wr.NewChoice("Strike", 6),
		wr.NewChoice("Ball In Play", 4),
	)

	for p := 1; p <= 30; p++ {

		pitch := chooser.Pick()

		switch {
		case pitch == "Ball":
			walk++
			fmt.Printf("Pitch %v:\tBall %v\n", p, walk)

			if walk == 4 {
				fmt.Println("That's a walk, take your base.")
				BaseTracking(g, b, 1)
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
					g.HalfInning++
					fmt.Println()
					fmt.Println()
				}
				return
			}
		case pitch == "Ball In Play":
			fmt.Printf("Pitch %v:\t", p)
			BallInPlay(g, b)
			fmt.Println()
			fmt.Println()
			return

		}
		time.Sleep(pause)
	}

}

func BallInPlay(g *Game, b *Bases) {
	// [x] single
	// [x] double
	// [x] triple
	// [x] homerun
	// [] account for people on base
	// [] keep track of runs scored
	// [x] catcher's interference
	// [x] hit that doesn't cross in play line
	// [x] hit caught by pitcher
	// [x] hit caught by infielder (2)

	rand.Seed(time.Now().UnixNano())

	result := rand.Intn(9) + 1

	switch result {
	case 1:
		fmt.Print("Single\n")
		BaseTracking(g, b, 1)
	case 2:
		fmt.Printf("Double\n")
		BaseTracking(g, b, 2)
	case 3:
		fmt.Printf("Triple\n")
		BaseTracking(g, b, 3)
	case 4:
		fmt.Printf("Homerun\n")
		BaseTracking(g, b, 4)
	case 5:
		fmt.Printf("catcher's interference\n")
		BaseTracking(g, b, 1)
	case 6:
		fmt.Printf("ground out to pitcher\n")
		g.Outs++
		if g.Outs == 3 {
			g.HalfInning++
		}
	case 7:
		fmt.Printf("ground out to infielder\n")
		g.Outs++
		if g.Outs == 3 {
			g.HalfInning++
		}
	case 8:
		fmt.Printf("Flyout to outfielder\n")
		g.Outs++
		if g.Outs == 3 {
			g.HalfInning++
		}
	case 9:
		fmt.Printf("Hit too softly, that's an out\n")
		g.Outs++
		if g.Outs == 3 {
			g.HalfInning++
		}
	default:
		fmt.Printf("random number %v\n", result)
	}
}
