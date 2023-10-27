package main

import (
	game "ballSim/internal"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/inancgumus/screen"
)

type Teams struct {
	City    string
	Name    string
	Players []Player
}

type Player struct {
	Name      string
	JerseyNum int32
	BatsHand  string
	ThrowHand string
	Position  string
	Rating    int32
}

//  This function just clears the screen and moves the cursor to the top left.
func Clear() {
	// Clear the screen TODO: put this in a function, if I find I am doing this often
	screen.Clear()
	screen.MoveTopLeft()
}

func main() {

	rand.Seed(time.Now().UnixNano())

	// s := Teams{
	// 	City: "St. Louis",
	// 	Name: "Sluggers",
	// 	Players: []Player{
	// 		{
	// 			Name:      "Mike Dix",
	// 			JerseyNum: 7,
	// 			BatsHand:  "Right",
	// 			ThrowHand: "Right",
	// 			Position:  "SS",
	// 			Rating:    8,
	// 		}, {
	// 			Name:      "Avery Dix",
	// 			JerseyNum: 10,
	// 			BatsHand:  "Right",
	// 			ThrowHand: "Right",
	// 			Position:  "P",
	// 			Rating:    8,
	// 		},
	// 	},
	// }

	// p1 := player{"Avery", 6, "Right", "Right", 6, "St. Louis Sluggers"}
	// p2 := player{"Mike", 6, "Right", "Right", 8, "St. Louis Sluggers"}
	// p3 := player{"Grace", 6, "Right", "Right", 7, "Memphis Mafia"}
	// p4 := player{"Caleb", 6, "Right", "Right", 8, "Memphis Mafia"}

	// Start up screen flow
	Clear()
	fmt.Println("WELCOME TO MAJOR LEAGUE FUZZBALL")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	pause := 3 * time.Second
	time.Sleep(pause)

	fmt.Printf("\tPress \"ENTER\" key to begin...\n")
	fmt.Scanln()

	Clear()

	// Team selection screens
	fmt.Println("Select your team")
	fmt.Println()
	fmt.Printf("\t1.\tSt. Louis Sluggers\n")
	fmt.Printf("\t2.\tMemphis Mafia\n")

	// Team selection accepting user input
	var number int64
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		log.Fatal(err)
	}

	var playerTeam, cpuTeam string

	switch {
	case number == 1:
		fmt.Println("So you want to be a slugger huh?")
		playerTeam = "St. Louis"
		cpuTeam = "Memphis"
		//w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		//fmt.Println("Name\t\tNumber\tBats\tThrows\tPosition\tRating")
		// Open the File
		csvFile, err := os.Open("../players.csv")
		if err != nil {
			fmt.Println(err)
		}

		// Read the File
		r := csv.NewReader(csvFile)

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(record)

		}
	// Save File to Struct
	case number == 2:
		fmt.Println("You are a natural hitman!")
		playerTeam = "Memphis Mafia"
		cpuTeam = "St. Louis Sluggers"
	default:
		fmt.Println("Let's try that again")
	}

	time.Sleep(pause)

	Clear()

	fmt.Printf("Alright top of the first and %s will be looking to jump ahead\n", cpuTeam)
	fmt.Printf("Here we go, as %s takes the field\n", playerTeam)

	// Open the File
	csvFile, err := os.Open("../players.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	// Read the File
	r := csv.NewReader(csvFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(record)

	}
	Clear()
	// Save File to Struct

	game.PlayBall()
}
