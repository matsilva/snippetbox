package models

import (
	"errors"
	"time"
)

//ErrNoRecord is the default error when a record is not found
//This helps remove dependencies from from the underlying datastore
var ErrNoRecord = errors.New("models: no matching record found")

//Snippet contains the info for the snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
