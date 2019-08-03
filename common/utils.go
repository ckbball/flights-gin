package common

import ()

type CommonError struct {
  Errors map[string]interface{} `json:"errors"`
}

func NewError(key string, err error) CommonError {
  response := CommonError{}
  response.Errors = make(map[string]interface{})
  response.Errors[key] = err.Error()
  return response
}
