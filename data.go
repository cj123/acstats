package main

type RaceData struct {
	Extras           []RaceExtras `json:"extras"`
	NumberOfSessions int64        `json:"number_of_sessions"`
	Players          []Player     `json:"players"`
	Sessions         []Session    `json:"sessions"`
	Track            string       `json:"track"`
}

type Session struct {
	BestLaps  []BestLap `json:"bestLaps"`
	Duration  int64     `json:"duration"`
	Event     int64     `json:"event"`
	Laps      []Lap     `json:"laps"`
	LapsCount int64     `json:"lapsCount"`
	Lapstotal []int64   `json:"lapstotal"`
	Name      string    `json:"name"`
	Type      int64     `json:"type"`
}

type Lap struct {
	Car     int64   `json:"car"`
	Cuts    int64   `json:"cuts"`
	Lap     int64   `json:"lap"`
	Sectors []int64 `json:"sectors"`
	Time    int64   `json:"time"`
	Tyre    string  `json:"tyre"`
}

type BestLap struct {
	Car  int64 `json:"car"`
	Lap  int64 `json:"lap"`
	Time int64 `json:"time"`
}

type Player struct {
	Car  string `json:"car"`
	Name string `json:"name"`
	Skin string `json:"skin"`
}

type RaceExtras struct {
	Name string `json:"name"`
	Time int64  `json:"time"`
}
