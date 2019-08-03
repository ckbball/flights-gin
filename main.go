package main

import (
  //"database/sql"
  //"fmt"
  // "database/sql"
  "fmt"
  "github.com/ckbball/flight-gin/common"
  "github.com/ckbball/flight-gin/flights"
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
)

/*
func Migrate(db *sql.DB) {
  db.AutoMigrate(&flights.FlightModel{})
}*/

func main() {

  _, err := common.Init()
  if err != nil {
    fmt.Println(err)
  }
  // Migrate(db.db)
  defer common.Close()

  r := gin.Default()

  v1 := r.Group("/api")
  flights.FlightsRegister(v1.Group("/flights"))

  r.Run()

}
