package db

import (
	"database/sql"
	"server/pj"

	_ "github.com/go-sql-driver/mysql"
)

type MultInit struct{}

func (MultInit) Error() string {
	return "Multiple initializations"
}

var (
	db *sql.DB
	mi MultInit
)

func InitDB() error {
	if db != nil {
		return mi
	}

	tmp, err := sql.Open("mysql", "root:12345678@tcp(47.120.8.50:3306)/FundsAndIndices?parseTime=true")
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

func CHRecordQuery(sd string, ed string, symbols []string) ([][]interface{}, error) {
	var RawData [][]interface{}
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

	var tmp = make([]interface{}, 4)
	for rows.Next() {
		if err := rows.Scan(&tmp[0], &tmp[1], &tmp[2], &tmp[3]); err != nil {
			continue
		}
		RawData = append(RawData, tmp)
	}

	return RawData, nil
}
