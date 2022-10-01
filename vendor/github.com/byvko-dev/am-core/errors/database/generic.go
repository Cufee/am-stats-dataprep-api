package database

import "errors"

var ErrDocumentNotFound = errors.New("not found")
var ErrDocumentExists = errors.New("document exists")
var ErrCollectionDoesNotExist = errors.New("collection does not exist")
