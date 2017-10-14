package main

import(
 "github.com/labstack/echo"
 "net/http"
 "strconv"
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

func getUser(id int) *User {
  email := strconv.Itoa(id) + "@example.com"
  user := &User {
    Id:    id,
    Email: email,
  }
  return user
}
