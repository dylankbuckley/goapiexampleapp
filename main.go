package main

import (
    // "fmt"
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode("Hello World, My Name is Dylan. This is my first GO app")
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", HelloWorld).Methods("GET")
    log.Fatal(http.ListenAndServe(GetRuntimePort(), router))
}

func GetRuntimePort() string {
  if len(os.Getenv("PORT")) > 0 {
    return ":"+os.Getenv("PORT")
  } else {
    return ":8080"
  }
}
