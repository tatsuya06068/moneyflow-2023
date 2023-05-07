package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	handler2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello-2\n")
	}

	// パスとハンドラー関数を結びつける
	mux.HandleFunc("/foo/", h)
	mux.HandleFunc("/", handler2)
	// mux.HandleFunc("/data/", GetData)

	// サーバを起動
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func h(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello-1***\n")
}

type Category struct {
	ID   int    `db:"m_bop_category_id"`
	Name string `db:"bop_name"`
}

type Json struct {
	Status int        `json:"status"`
	Result []Category `json:"category"`
}

// func GetData(w http.ResponseWriter, _ *http.Request) {

// 	// JWTに付与する構造体
// 	claims := jwt.MapClaims{
// 		"user_id": "user_id1234",
// 		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24時間が有効期限
// 	}

// 	// ヘッダーとペイロード生成
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// トークンに署名を付与
// 	accessToken, _ := token.SignedString([]byte("ACCESS_SECRET_KEY"))
// 	fmt.Println("accessToken:", accessToken)

// 	dbDriver := driver.NewSqlHandler()

// 	row, err := dbDriver.Query("select m_bop_category_id, bop_name from m_bop_categories")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer row.Close()

// 	ca := []Category{}
// 	for row.Next() {
// 		c := &Category{}
// 		err := row.Scan(&c.ID, &c.Name)
// 		if err != nil {
// 			panic(err)
// 		}
// 		ca = append(ca, *c)
// 	}
// 	j := Json{Status: 200, Result: ca}
// 	res, _ := json.Marshal(j)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(res)
// }
