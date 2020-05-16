package model

import (
	"errors"
)

var (
	InvalidRequest = errors.New("Error:: Invalid Request")
	DBErrNotFound  = errors.New("Error:: record not exists")
)
