package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", HelloServer)

    httpPort := os.Getenv("HTTP_PORT")
    if httpPort == "" {
      httpPort = "8080"
    }

    http.ListenAndServe(":" + httpPort, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    name, err := os.Hostname()
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w, "Hello, from  %s!", name)
}
