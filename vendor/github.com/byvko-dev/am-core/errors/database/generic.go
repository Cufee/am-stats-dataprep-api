package database

import "errors"

var DocumentNotFound = errors.New("not found")
var DocumentExists = errors.New("document exists")
