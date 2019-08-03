package flights

import (
  //"errors"
  //"encoding/json"
  "fmt"
  "github.com/ckbball/flight-gin/common"
  "github.com/gin-gonic/gin"
  "net/http"
  //"strconv"
)

func FlightsRegister(router *gin.RouterGroup) {
  router.POST("/add", FlightCreate)
  //router.GET("/flights/:id", FlightsGet)
  router.GET("", FlightsGetAll)
}

func FlightGet(c *gin.Context) {

}

func FlightsGetAll(c *gin.Context) {

  f, err := GetAllFlights()
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
  }
  fmt.Println("flights from FlightsGetAll router.go")
  fmt.Println(f)

  serializer := FlightsSerializer{c, f}
  c.JSON(http.StatusOK, gin.H{"flights": serializer.Response()})
}

func FlightCreate(c *gin.Context) {
  var v = NewFlightModelValidator()

  c.BindJSON(&v)
  fmt.Println("Req body")
  fmt.Println(v.Flight.DepartureCity)

  flight, _ := FlightValidatorToModel(v)

  if err := SaveFlight(flight); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }

  serializer := FlightSerializer{c, flight}
  c.JSON(http.StatusCreated, gin.H{"flight": serializer.Response()})
}
