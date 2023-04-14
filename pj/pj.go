package pj

import "time"

type Indice struct {
	Symbol string    `json:"symbol"`
	Name   string    `json:"name"`
	EDate  time.Time `json:"earliest_date"`
}

type IndicesData struct {
	Indices []Indice `json:"indices"`
}
