package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
	"strconv" // 文字列を数値に変換
)

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type Users []User

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

	e.GET("/users", func(c echo.Context) error {
		userList()
		return c.JSON(http.StatusOK, getUser(1))
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func userList() {
	db, err := sql.Open("mysql", "root:@/famiphotos_development")
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

	var users Users

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		user := User{}
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)

			if columns[i] == "Email" {
				user.Email = value
			}
		}
		users = append(users, user)
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
