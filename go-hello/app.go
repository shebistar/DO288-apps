package main

import (
    "net/http"
    "flag"
    "fmt"
)

var lang = flag.String("lang", "en", "run app with language support - default is english")

func main() {
    var port = "8080"
    http.HandleFunc("/", rootHandler)
    fmt.Printf("Starting server on port %v...\n", port)
    http.ListenAndServe(":"+port, nil)
}

func rootHandler(response http.ResponseWriter, request *http.Request) {

    flag.Parse()

    switch *lang {
      case "en":
        fmt.Fprintf(response, "Helloooooo %s!. Welcomeeeeeeeeeee Quiero que esto esté en español, vamos a ver si se puede!\n", request.URL.Path[1:])
      case "es":
        fmt.Fprintf(response, "Hola %s!. Bienvenido!\n", request.URL.Path[1:])
      default:
        fmt.Fprintf(response, "Error! unknown lang option -> %s\n", *lang)
  }
}
