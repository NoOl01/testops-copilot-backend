package consts

import "errors"

const (
	BadRequest = "bad request"
)

var (
	GenerateInvalidBody = errors.New("invalid request body")
	ConnectionTimeout   = errors.New("connection timeout")
	RequestCanceled     = errors.New("request canceled")
)
