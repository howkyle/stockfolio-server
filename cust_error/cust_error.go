package cust_error

import "errors"

// "net/http"

var NotFound = errors.New("not found")

// type NotFound struct {
// 	Err          error
// 	ResponseCode int
// 	Message      string
// }

// func (n NotFound) Error() string {
// 	return fmt.Sprintf("%s: %v", n.Message, n.Err.Error())
// }

// func (n NotFound) Code() int {
// 	return n.ResponseCode
// }

// func NewNotFound(message string, err error) error {
// 	return NotFound{Err: err, ResponseCode: http.StatusNotFound, Message: message}
// }
