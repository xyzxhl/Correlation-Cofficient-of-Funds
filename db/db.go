package db

import (
	"database/sql"
	"server/pj"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MultInit struct{}

func (MultInit) Error() string {
	return "Multiple initializations"
}

type Record struct {
	Date    time.Time
	Percent float32
}

var (
	db *sql.DB
	mi MultInit
)

func InitDB() error {
	if db != nil {
		return mi
	}

	tmp, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/Funds_Indices?parseTime=true")
	if err != nil {
		return err
	}

	db = tmp
	return nil
}

func FINameGetAll() (pj.IndicesData, error) {
	var IndicesData pj.IndicesData
	rows, err := db.Query("SELECT * FROM FIName")
	if err != nil {
		return IndicesData, err
	}

	var tmp pj.Indice
	for rows.Next() {
		if err := rows.Scan(&tmp.Symbol, &tmp.Name, &tmp.EDate); err != nil {
			continue
		}
		IndicesData.Indices = append(IndicesData.Indices, tmp)
	}

	return IndicesData, nil
}

func CHRecordQuery(sd string, ed string, symbols []string) (map[string][]Record, error) {
	var RawData = make(map[string][]Record)
	inst := "SELECT * FROM CHRecord WHERE date BETWEEN '" + sd + "' AND '" + ed + "' AND symbol IN ('"
	for i, symbol := range symbols {
		if i != len(symbols)-1 {
			inst += symbol + "','"
		} else {
			inst += symbol + "')"
		}
	}

	rows, err := db.Query(inst)
	if err != nil {
		return RawData, err
	}

	var tmp Record
	var sym string
	for rows.Next() {
		if err := rows.Scan(&sym, &tmp.Date, &tmp.Percent); err != nil {
			continue
		}
		RawData[sym] = append(RawData[sym], tmp)
	}

	return RawData, nil
}
