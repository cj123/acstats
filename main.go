package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
	"time"
)

type SessionData struct {
	LapNum         int64
	Time           time.Duration
	NumCuts        int64
	AverageSoFar   time.Duration
	DeltaOnAvg     time.Duration
	BestLap        bool
	LikelyCheating bool
}

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "filename of json data")
	flag.Parse()

	if filename == "" {
		fmt.Printf("Specify a filename, e.g. %s -f race-data.json\n", os.Args[0])
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var raceData RaceData

	err = json.Unmarshal(data, &raceData)

	if err != nil {
		panic(err)
	}

	for playerIndex, player := range raceData.Players {
		if player.Name == "" {
			continue // empty slots
		}

		fmt.Printf("Player: %s\n", player.Name)

		for _, session := range raceData.Sessions {
			if session.Lapstotal[playerIndex] < 1 {
				continue
			}

			sessionData := make([]SessionData, session.Lapstotal[playerIndex])

			fmt.Printf("Session: %s (%d laps)\n\n", session.Name, session.Lapstotal[playerIndex])

			var bestLap BestLap

			for _, lap := range session.BestLaps {
				if lap.Car != int64(playerIndex) {
					continue
				}

				bestLap = lap
			}

			for _, lap := range session.Laps {
				if lap.Car != int64(playerIndex) {
					continue
				}

				avg := int64(session.playerAverageUntil(playerIndex, lap.Lap))
				sessionData[lap.Lap] = SessionData{
					LapNum:         lap.Lap,
					Time:           time.Duration(lap.Time) * time.Millisecond,
					NumCuts:        lap.Cuts,
					AverageSoFar:   time.Duration(avg) * time.Millisecond,
					DeltaOnAvg:     time.Duration(lap.Time-avg) * time.Millisecond,
					BestLap:        lap.Lap == bestLap.Lap,
					LikelyCheating: lap.Time < avg && lap.Cuts > 0,
				}
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Lap", "Time", "Cuts", "Avg So Far", "Delta on Avg", "Best Lap?", "Cheating?"})
			table.SetBorder(false)

			for _, v := range sessionData {
				table.Append([]string{fmt.Sprintf("%d", v.LapNum), v.Time.String(), fmt.Sprintf("%d", v.NumCuts), v.AverageSoFar.String(), v.DeltaOnAvg.String(), boolToString(v.BestLap), boolToString(v.LikelyCheating)})
			}

			table.Render()
			fmt.Println()
		}

		fmt.Printf("\n\n * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *\n\n")
	}
}

func boolToString(x bool) string {
	if x {
		return "yes"
	} else {
		return "no"
	}
}

func (s *Session) playerAverageUntil(playerIndex int, lapIndex int64) float64 {
	var sum int64 = 0

	for _, lap := range s.Laps {
		if lap.Car != int64(playerIndex) {
			continue
		}

		if lap.Lap > lapIndex {
			break
		}

		sum += lap.Time
	}

	return float64(sum) / float64(lapIndex+1)
}
