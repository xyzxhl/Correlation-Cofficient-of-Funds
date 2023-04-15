package pj

import (
	"time"
)

type Indice struct {
	Symbol string    `json:"symbol"`
	Name   string    `json:"name"`
	EDate  time.Time `json:"earliest_date"`
}

type IndicesData struct {
	Indices []Indice `json:"indices"`
}

type DailyChange struct {
	Date   time.Time `json:"date"`
	Change float32   `json:"change"`
}

type IndiceChanges struct {
	Symbol  string        `json:"symbol"`
	Changes []DailyChange `json:"changes"`
}

type IndicesChangesData struct {
	IndicesCh []IndiceChanges `json:"indices_change"`
}
