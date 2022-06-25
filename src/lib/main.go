package jd

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Aesthetic struct {
	Name        string
	Description string
	Outcomes    string
}

type Entry struct {
	Datetime  string
	Gratitude string
	Note      string
}

type Step struct {
	Tracked string // The what you are tracking
	Grade   int    // The grade you gave yourself
}

type Daily struct {
	Datetime string
	Ratings  []Step
}

//const dbFile string = "/tmp/journal-dandy.db"

func AddAest(name string, desc string, outc string) Aesthetic {
	aest := Aesthetic{
		Name:        name,
		Description: desc,
		Outcomes:    outc,
	}
	return aest
}
func AddEntry(grat string, note string) Entry {
	entry := Entry{
		Datetime:  time.Now().Format(time.RFC3339),
		Gratitude: grat,
		Note:      note,
	}
	return entry
}
func AddStep(trac string, grade int) Step {
	step := Step{
		Tracked: trac,
		Grade:   grade,
	}
	return step
}
func AddDaily(ratings []Step) Daily {
	daily := Daily{
		Datetime: time.Now().Format(time.RFC3339),
		Ratings:  ratings,
	}
	return daily
}
