package app

import "errors"

var ErrNotFound = errors.New("Not Found")
var ErrWrongCredentials = errors.New("Username or password is incorrect")
