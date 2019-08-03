package flights

type FlightModelValidator struct {
  Flight struct {
    DepartureCity string `json:"departurecity"`
    ArrivalCity   string `json:"arrivalcity"`
    Airline       string `json:"airline"`
    AirlineID     uint   `json:"airlineid"`
    DepartAt      string `json:"departat"`
    ArriveAt      string `json:"arriveat"`
  } `json:"flight"`
  flightModel FlightModel `json:"-"`
}

func NewFlightModelValidator() FlightModelValidator {
  return FlightModelValidator{}
}
