package flights

import (
  //"errors"
  //"encoding/json"
  "fmt"
  "github.com/ckbball/flight-gin/common"
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"
  "strings"
  "unicode"
)

func FlightsRegister(router *gin.RouterGroup) {
  router.POST("/add", FlightCreate)
  router.GET("/:id", FlightGet)
  router.GET("", FlightsGetAll)
  router.GET("/", FlightsFiltered)
}

// ---------------------- ROUTER FUNCTIONS ----------------------------

func FlightsFiltered(c *gin.Context) {
  departurecity := c.Query("dc")
  arrivalcity := c.Query("ac")
  departat := c.Query("da")

  f, err := FilteredFlights(departurecity, arrivalcity, departat)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
  }
  fmt.Println("flights matching query: %v, %v, %v from FlightsFiltered router.go", departurecity, arrivalcity, departat)
  fmt.Println(f)

  serializer := FlightsSerializer{c, f}
  c.JSON(http.StatusOK, gin.H{"flights": serializer.Response()})

}

func FlightGet(c *gin.Context) {
  id := c.Param("id")
  Id, err := strconv.Atoi(id)

  f, err := GetFlight(Id)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
  }
  fmt.Println("flight of id: %v from FlightGet router.go", f.ID)
  fmt.Println(*f)

  serializer := FlightSerializer{c, *f}
  c.JSON(http.StatusOK, gin.H{"flight": serializer.Response()})
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

  vs := removeSpace(v)

  flight, _ := FlightValidatorToModel(vs)

  if err := SaveFlight(flight); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }

  serializer := FlightSerializer{c, flight}
  c.JSON(http.StatusCreated, gin.H{"flight": serializer.Response()})
}

func removeSpace(v FlightModelValidator) FlightModelValidator {
  v.Flight.DepartureCity = stripSpaces(v.Flight.DepartureCity)
  v.Flight.ArrivalCity = stripSpaces(v.Flight.ArrivalCity)
  v.Flight.Airline = stripSpaces(v.Flight.Airline)
  return v
}

func stripSpaces(str string) string {
  return strings.Map(func(r rune) rune {
    if unicode.IsSpace(r) {
      return -1
    }
    return r
  }, str)
}
