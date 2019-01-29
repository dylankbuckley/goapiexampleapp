package main

import (
    "fmt"
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
    "github.com/globalsign/mgo/bson"
    "strings"
)

type (
  Job struct {
    EmployerName string `bson:"employerName"`
    JobTitle string `bson:"jobTitle"`
    LocationName string `bson:"locationName"`
    JobDescription string `bson:"jobDescription"`
    Applications int32 `bson:"applications"`
    Date string `bson:"date"`
    MinimumSalary int32 `bson:"minimumSalary"`
    MaximumSalary int32 `bson:"maximumSalary"`
    Currency string `bson:"currency"`
    DirectEmployer bool `bson:"directEmployer"`
    JobURL string `bson:"jobUrl"`
    ExpirationDate string `bson:"expirationDate"`
  }
  QueryUniverse struct {
    Results int32 `bson:"results"`
  }
  QueryCleaned struct {
    Cleaned []Job `bson:"cleaned"`
  }
)

func main() {
  // Connect to MongoDB & Create Server
  connectionFailure := connectToMongoDb()
  if connectionFailure == true {
    fmt.Println("Failure Connecting to MongoDB")
  } else {
    fmt.Println("Success Connecting to MongoDB")
    // Create Web Server & Endpoints
    router := mux.NewRouter()
    //router.HandleFunc("/", TotalQueryUniverse).Methods("GET")
    router.HandleFunc("/{term}/{location}", TotalQueryUniverse).Methods("GET")
    router.HandleFunc("/{term}/{location}/d", TotalQueryCleaned).Methods("GET")
    log.Fatal(http.ListenAndServe(GetRuntimePort(), router))
  }
}

func TotalQueryUniverse(w http.ResponseWriter, r *http.Request) {
  // Get Query Parameters
  params := mux.Vars(r)
  // Clone Session
  session := mongoSession.Clone()
  // Connect to Collection
  collection := session.DB("counterjobs").C("queries")
  var queries []QueryUniverse
  // Perform Query
  err1 := collection.Find(bson.M{"location": strings.Replace(params["location"], "-", " ", -1), "query": strings.Replace(params["term"], "-", " ", -1)}).All(&queries)
  if err1 != nil {
    fmt.Println(err1, "Error")
  }
  // Close Session
  session.Close()
  // Send Response to Client
  json.NewEncoder(w).Encode(queries[0])
}

func TotalQueryCleaned(w http.ResponseWriter, r *http.Request) {
  // Get Query Parameters
  params := mux.Vars(r)
  // Clone Session
  session := mongoSession.Clone()
  // Connect to Collection
  collection := session.DB("counterjobs").C("queries")
  var queries []QueryCleaned
  // Perform Query
  err1 := collection.Find(bson.M{"location": strings.Replace(params["location"], "-", " ", -1), "query": strings.Replace(params["term"], "-", " ", -1)}).All(&queries)
  if err1 != nil {
    fmt.Println(err1, "Error")
  }
  // Close Session
  session.Close()
  // Send Response to Client
  json.NewEncoder(w).Encode(queries[0])
}


func GetRuntimePort() string {
  if len(os.Getenv("PORT")) > 0 {
    return ":"+os.Getenv("PORT")
  } else {
    return ":8080"
  }
}
