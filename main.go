package main

import (
	"./handler"
	"github.com/labstack/echo"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()
	// ルーティング
	// e.GET("/hello", handler.MainPage())
	e.GET("/hello/:username", handler.MainPage())
	// サーバー起動
	e.Start(":1323") //ポート番号指定してね
}
