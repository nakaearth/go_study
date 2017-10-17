package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
	"strconv" // 文字列を数値に変換
)

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()

	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, getUser(id))
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func getUserList() {
	db, err := sql.Open("mysql", "root:@/my_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM users") //
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
}

func getUser(id int) *User {
	email := strconv.Itoa(id) + "@example.com"
	user := &User{
		Id:    id,
		Email: email,
	}
	return user
}
