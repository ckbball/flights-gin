package flights

import (
  "github.com/gin-gonic/gin"
)

type FlightResponse struct {
  DepartureCity string `json:"departurecity"`
  ArrivalCity   string `json:"arrivalcity"`
  Airline       string `json:"airline"`
  AirlineID     uint   `json:"airlineid"`
  DepartAt      string `json:"departat"`
  ArriveAt      string `json:"arriveat"`
}

type FlightSerializer struct {
  C *gin.Context
  FlightModel
}

type FlightsSerializer struct {
  C  *gin.Context
  Fs []*FlightModel
}

func (s *FlightsSerializer) Response() []*FlightResponse {
  var response []*FlightResponse

  for i := 0; i < len(s.Fs); i++ {
    flight := &FlightResponse{
      DepartureCity: s.Fs[i].DepartureCity,
      ArrivalCity:   s.Fs[i].ArrivalCity,
      Airline:       s.Fs[i].Airline,
      AirlineID:     s.Fs[i].AirlineID,
      DepartAt:      s.Fs[i].DepartAt,
      ArriveAt:      s.Fs[i].ArriveAt,
    }

    response = append(response, flight)
  }

  return response
}

func (s *FlightSerializer) Response() FlightResponse {
  response := FlightResponse{
    DepartureCity: s.DepartureCity,
    ArrivalCity:   s.ArrivalCity,
    Airline:       s.Airline,
    AirlineID:     s.AirlineID,
    DepartAt:      s.DepartAt,
    ArriveAt:      s.ArriveAt,
  }
  return response
}
