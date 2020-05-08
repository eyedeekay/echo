package main

import (
	"github.com/eyedeekay/echo/handler"
	"net/http"
)

func main() {
	echobot, _ := echo.NewEchoBot()

	http.ListenAndServe(":7777", echobot)
}
