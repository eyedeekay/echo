package echo

import (
    "net/http"
)

type Echo struct {
    *http.ServeMux
}

func NewEchoBot() (*Echo, error) {
    var e Echo
    e.ServeMux = http.NewServeMux()
    e.ServeMux.HandleFunc("/", EchoHandler)
    return &e, nil
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
    for key, value := range r.Header {
        w.Write([]byte(key + "="))
        for _, value := range value {
            w.Write([]byte(value))
        }
        w.Write([]byte("\n"))
    }
}