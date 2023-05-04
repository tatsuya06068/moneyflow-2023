package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/tatsuya06068/moneyflow-2023/driver"
)

func main() {

	handler2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello-2\n")
	}

	// パスとハンドラー関数を結びつける
	http.HandleFunc("/foo/", h)
	http.HandleFunc("/bar/", handler2)
	http.HandleFunc("/data/", GetData)

	// サーバを起動
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func h(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello-1\n")
}

type Category struct {
	ID   int    `db:"m_bop_category_id"`
	Name string `db:"bop_name"`
}

type Json struct {
	Status int        `json:"status"`
	Result []Category `json:"category"`
}

func GetData(w http.ResponseWriter, _ *http.Request) {
	dbDriver := driver.NewSqlHandler()

	row, err := dbDriver.Query("select m_bop_category_id, bop_name from m_bop_categories")
	if err != nil {
		panic(err)
	}

	defer row.Close()

	ca := []Category{}
	for row.Next() {
		c := &Category{}
		err := row.Scan(&c.ID, &c.Name)
		if err != nil {
			panic(err)
		}
		ca = append(ca, *c)
	}
	j := Json{Status: 200, Result: ca}
	res, _ := json.Marshal(j)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}
