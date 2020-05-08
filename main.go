package main

import (
    "net/http"
    "github.com/eyedeekay/echo/handler"
)

func main(){
    echobot, _ := echo.NewEchoBot()

    http.ListenAndServe(":7777", echobot)
}