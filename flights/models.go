package flights

import (
  //"errors"
  "fmt"
  "github.com/ckbball/flight-gin/common"
  //"github.com/gin-gonic/gin/binding"
  "database/sql"
  //"strconv"
  //"time"
)

type FlightModel struct {
  ID            int
  DepartureCity string `json:"departurecity"`
  ArrivalCity   string `json:"arrivalcity"`
  Airline       string `json:"airline"`
  AirlineID     uint   `json:"airlineid"`
  DepartAt      string `json:"departtime"`
  ArriveAt      string `json:"arriveat"`
}

/*
func GetFlightsByAirline(id uint) ([]FlightModel, error) {

  db := common.GetDB()
  flights := make([]FlightModel, 0)
  tx := db.Begin()
  tx.Where("ID = ?", id).Find(flights)
  err := tx.Commit().Error
  return flights, err
}*/

// -------- HELPER FUNCTIONS BEGIN --------------------------------------

func FlightValidatorToModel(newFlight FlightModelValidator) (FlightModel, error) {
  f := FlightModel{}

  f.DepartureCity = newFlight.Flight.DepartureCity
  f.ArrivalCity = newFlight.Flight.ArrivalCity
  f.Airline = newFlight.Flight.Airline
  f.AirlineID = newFlight.Flight.AirlineID
  f.DepartAt = newFlight.Flight.DepartAt
  f.ArriveAt = newFlight.Flight.ArriveAt

  return f, nil
}

func scanFlight(s *sql.Rows) (*FlightModel, error) {
  var (
    ID            int
    DepartureCity sql.NullString
    ArrivalCity   sql.NullString
    Airline       sql.NullString
    AirlineID     uint
    DepartAt      sql.NullString
    ArriveAt      sql.NullString
  )

  if err := s.Scan(&ID, &DepartureCity, &ArrivalCity, &Airline, &AirlineID, &DepartAt, &ArriveAt); err != nil {
    return nil, err
  }

  id := int(ID)

  flight := &FlightModel{
    ID:            id,
    DepartureCity: DepartureCity.String,
    ArrivalCity:   ArrivalCity.String,
    Airline:       Airline.String,
    AirlineID:     AirlineID,
    DepartAt:      DepartAt.String,
    ArriveAt:      ArriveAt.String,
  }

  return flight, nil
}

// ------------- HELPER FUNCTIONS END ----------------------------

// --------------- DB FUNCTIONS BEGIN -------------------------------------

func GetAllFlights() ([]*FlightModel, error) {
  rows, err := common.GetRows("flights")
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  var flights []*FlightModel
  for rows.Next() {
    flight, err := scanFlight(rows)
    if err != nil {
      return nil, fmt.Errorf("ERROR --> mysql: could not read row: %v", err)
    }

    flights = append(flights, flight)
  }

  return flights, nil
}

func SaveFlight(f FlightModel) error {

  fmt.Println("Finished copying validator to model for db creation")

  r, err := common.ExecAffectingOneRow("insert", f.DepartureCity, f.ArrivalCity, f.Airline, f.AirlineID, f.DepartAt, f.ArriveAt)
  if err != nil {
    return err
  }
  fmt.Println("Insert flight result: %v", r)

  return nil
}

// ----------------- DB FUNCTIONS END ------------------------------
