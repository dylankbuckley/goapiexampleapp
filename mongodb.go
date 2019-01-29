package main

import (
  "crypto/tls"
  "github.com/globalsign/mgo"
  "net"
)

var mongoSession *mgo.Session

// Connect to MongoDB & Return Session
func connectToMongoDb() (bool) {
  tlsConfig := &tls.Config{}
  dialInfo := &mgo.DialInfo{
      Addrs: []string{"counter-cms-shard-00-00-izr7l.mongodb.net:27017",
                      "counter-cms-shard-00-01-izr7l.mongodb.net:27017",
                      "counter-cms-shard-00-02-izr7l.mongodb.net:27017"},
      Database: "admin",
      Username: "gobot",
      Password: "gobot123",
  }
  dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
      conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
      return conn, err
  }
  session, err := mgo.DialWithInfo(dialInfo)
  if err != nil {
    return true
  }
  mongoSession = session
  return false
}
